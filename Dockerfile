FROM debian:stretch

WORKDIR /home/homo

# Install system dependence
RUN \
    apt-get update && \
    apt-get install -y git gcc automake autoconf libtool build-essential && \
    apt-get install -y bison swig python-dev libpulse-dev portaudio19-dev

# Install sphinxbase
RUN \
    git clone https://github.com/cmusphinx/sphinxbase.git && \
    cd sphinxbase && ./autogen.sh && ./configure && make -j 4 && make install && \
    cd .. && rm -rf sphinxbase

# Install PocketSphinx
RUN \
    git clone https://github.com/cmusphinx/pocketsphinx.git && \
    cp -r pocketsphinx/model/en-us sphinx/ && \
    cd pocketsphinx && ./autogen.sh && ./configure && make -j 4 && make install && \
    cd .. && rm -rf pocketsphinx

# Replace 1000 with your user / group id
#RUN export uid=1000 gid=1000 && \
#    mkdir -p /home/homo && \
#    echo "homo:x:${uid}:${gid}:homo,,,:/home/homo:/bin/bash" >> /etc/passwd && \
#    echo "homo:x:${uid}:" >> /etc/group && \
#    echo "homo ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/homo && \
#    chmod 0440 /etc/sudoers.d/homo && \
#    chown ${uid}:${gid} -R /home/homo

USER homo
ENV HOME /home/homo
