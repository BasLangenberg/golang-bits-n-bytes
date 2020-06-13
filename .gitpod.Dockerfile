FROM ninckblokje/golang-bits-n-bytes

RUN groupadd gitpod \
    && useradd -s /bin/bash -g gitpod -m gitpod

USER gitpod
WORKDIR /home/gitpod

# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN sudo apt-get -q update && #     sudo apt-get install -yq bastet && #     sudo rm -rf /var/lib/apt/lists/*
#
# More information: https://www.gitpod.io/docs/config-docker/

ENV PATH=${PATH}:/go/bin:/usr/local/go/bin

RUN mkdir -p /home/gitpod/.bashrc.d \
    && echo 'export PATH="${PATH}:/go/bin:/usr/local/go/bin:/usr/local/tinygo/bin"' >> /home/gitpod/.bashrc.d/env.sh
