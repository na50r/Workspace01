FROM jenkins/jenkins:lts

USER root

# Install necessary packages
RUN apt-get update && \
    apt-get install -y docker.io sudo python3 python3-pip wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    echo "jenkins ALL=NOPASSWD: ALL" >> /etc/sudoers

# Install Go 1.22.5
ENV GO_VERSION=1.22.5
RUN wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
    rm go$GO_VERSION.linux-amd64.tar.gz

# Add Go binary to PATH
ENV PATH=/usr/local/go/bin:$PATH

# Verify Go installation
RUN go version

# Switch back to Jenkins user
USER jenkins
