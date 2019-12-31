package master

import (
	"encoding/json"
	"fmt"
	"github.com/countstarlight/homo/logger"
	"github.com/countstarlight/homo/master/engine"
	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/countstarlight/homo/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"
)

type infoStats struct {
	homo.Inspect
	services engine.ServicesStats
	file     string
	sync.RWMutex
}

func newInfoStats(pwd, mode, version, revision, file string) *infoStats {
	return &infoStats{
		file:     file,
		services: engine.ServicesStats{},
		Inspect: homo.Inspect{
			Software: homo.Software{
				OS:          runtime.GOOS,
				Arch:        runtime.GOARCH,
				GoVersion:   runtime.Version(),
				PWD:         pwd,
				Mode:        mode,
				BinVersion:  version,
				GitRevision: revision,
			},
			Hardware: homo.Hardware{
				HostInfo: utils.GetHostInfo(),
				NetInfo:  utils.GetNetInfo(),
			},
		},
	}
}

func (is *infoStats) SetInstanceStats(serviceName, instanceName string, partialStats engine.PartialStats, persist bool) {
	is.Lock()
	service, ok := is.services[serviceName]
	if !ok {
		service = engine.InstancesStats{}
		is.services[serviceName] = service
	}
	instance, ok := service[instanceName]
	if !ok {
		instance = partialStats
		service[instanceName] = instance
	} else {
		for k, v := range partialStats {
			instance[k] = v
		}
	}
	if persist {
		is.persistStats()
	}
	is.Unlock()
}

func (is *infoStats) DelInstanceStats(serviceName, instanceName string, persist bool) {
	is.Lock()
	defer is.Unlock()
	service, ok := is.services[serviceName]
	if !ok {
		return
	}
	_, ok = service[instanceName]
	if !ok {
		return
	}
	delete(service, instanceName)
	if len(service) == 0 {
		delete(is.services, serviceName)
	}
	if persist {
		is.persistStats()
	}
}

func (is *infoStats) DelServiceStats(serviceName string, persist bool) {
	is.Lock()
	defer is.Unlock()
	_, ok := is.services[serviceName]
	if !ok {
		return
	}
	delete(is.services, serviceName)
	if persist {
		is.persistStats()
	}
}

func (is *infoStats) setVersion(ver string) {
	is.Lock()
	is.Inspect.Software.ConfVersion = ver
	is.Unlock()
}

func (is *infoStats) getVersion() string {
	is.RLock()
	defer is.RUnlock()
	return is.Inspect.Software.ConfVersion
}

func (is *infoStats) setError(err error) {
	is.Lock()
	if err == nil {
		is.Inspect.Error = ""
	} else {
		is.Inspect.Error = err.Error()
	}
	is.Unlock()
}

func (is *infoStats) getError() string {
	is.RLock()
	defer is.RUnlock()
	return is.Inspect.Error
}

// func genVolumesStats(cfg []homo.VolumeInfo) homo.Volumes {
// 	volumes := homo.Volumes{}
// 	for _, item := range cfg {
// 		volumes = append(volumes, homo.VolumeStatus{
// 			Name:    item.Name,
// 			Version: item.Meta.Version,
// 		})
// 	}
// 	return volumes
// }

func (is *infoStats) persistStats() {
	data, err := yaml.Marshal(is.services)
	if err != nil {
		logger.S.Warnf("failed to persist services stats: %s", err.Error())
		return
	}
	err = ioutil.WriteFile(is.file, data, 0755)
	if err != nil {
		logger.S.Warnf("failed to persist services stats: %s", err.Error())
	}
}

func (is *infoStats) LoadStats(services interface{}) bool {
	if !utils.IsFile(is.file) {
		return false
	}
	data, err := ioutil.ReadFile(is.file)
	if err != nil {
		logger.S.Warnf("failed to read old stats: %s", err.Error())
		os.Rename(is.file, fmt.Sprintf("%s.%d", is.file, time.Now().Unix()))
		return false
	}
	err = yaml.Unmarshal(data, services)
	if err != nil {
		logger.S.Warnf("failed to unmarshal old stats: %s", err.Error())
		os.Rename(is.file, fmt.Sprintf("%s.%d", is.file, time.Now().Unix()))
		return false
	}
	return true
}

func (is *infoStats) stats() {
	t := time.Now().UTC()
	gi := utils.GetGPUInfo()
	mi := utils.GetMemInfo()
	ci := utils.GetCPUInfo()
	di := utils.GetDiskInfo("/")

	is.Lock()
	is.Inspect.Time = t
	is.Inspect.Hardware.GPUInfo = gi
	is.Inspect.Hardware.MemInfo = mi
	is.Inspect.Hardware.CPUInfo = ci
	is.Inspect.Hardware.DiskInfo = di
	is.Unlock()
}

func (is *infoStats) serializeStats() ([]byte, error) {
	is.Lock()
	defer is.Unlock()

	result := is.Inspect
	result.Services = homo.Services{}
	for serviceName, serviceStats := range is.services {
		service := homo.NewServiceStatus(serviceName)
		for _, instanceStats := range serviceStats {
			service.Instances = append(service.Instances, map[string]interface{}(instanceStats))
		}
		result.Services = append(result.Services, service)
	}
	return json.Marshal(result)
}

// InspectSystem inspects info and stats of homo system
func (m *Master) InspectSystem() ([]byte, error) {
	defer logger.S.Debug("InspectSystem")
	var wg sync.WaitGroup
	for item := range m.services.IterBuffered() {
		wg.Add(1)
		go func(s engine.Service) {
			defer wg.Done()
			s.Stats()
		}(item.Val.(engine.Service))
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		m.infostats.stats()
	}()
	wg.Wait()

	return m.infostats.serializeStats()
}
