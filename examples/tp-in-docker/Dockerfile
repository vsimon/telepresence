# Demo Dockerfile for running telepresence from within a docker container.
# Run this image with:
#
#   docker run \
#      --privileged \
#      --device /dev/net/tun:/dev/net/tun \
#      --network=host \
#      -v /var/run/dbus/system_bus_socket:/var/run/dbus/system_bus_socket \
#      -v ~/.kube/config:/root/.kube/config \
#      -it --rm tp-in-docker /bin/bash
#
FROM alpine:3.13

# Install Telepresence prerequisites
RUN apk add --no-cache curl iproute2 sshfs bash

# Download and install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Get the telepresence binary from the download site
#  RUN curl -fL https://app.getambassador.io/download/tel2/linux/amd64/latest/telepresence -o telepresence
# or copy your own build (you probably want to combine this with --env TELEPRESENCE_REGISTRY=<your docker repo>)
COPY telepresence telepresence

# Install the telepresence binary
RUN install -o root -g root -m 0755 telepresence /usr/local/bin/telepresence

# The docker registry used when Telepresence installs the traffic-manager and traffic-agent
ENV TELEPRESENCE_REGISTRY=docker.io/datawire

# Add some conventient aliases to .bashrc
RUN echo -e '\
alias t=telepresence\n\
alias k=kubectl\n\
' >> /root/.bashrc

