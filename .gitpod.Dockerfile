FROM gitpod/workspace-full

USER root

RUN apt update \
        && apt install -y openjdk-11-jdk \
        && go get golang.org/x/tour \
        && wget https://github.com/tinygo-org/tinygo/releases/download/v0.13.1/tinygo_0.13.1_amd64.deb \
        && dpkg -i tinygo_0.13.1_amd64.deb
ENV PATH="/usr/local/tinygo/bin:/go/bin:${PATH}"

USER gitpod

# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN sudo apt-get -q update && #     sudo apt-get install -yq bastet && #     sudo rm -rf /var/lib/apt/lists/*
#
# More information: https://www.gitpod.io/docs/config-docker/
