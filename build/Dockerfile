# Step one: Build compliance-operator
FROM registry.access.redhat.com/ubi9/ubi:9.4-1181 AS builder

# Install Go 1.23 manually
RUN curl -O -L https://go.dev/dl/go1.23.0.linux-ppc64le.tar.gz && \
    tar -C /usr/local -xzf go1.23.0.linux-ppc64le.tar.gz && \
    rm go1.23.0.linux-ppc64le.tar.gz

# Install build tools including make
RUN dnf install -y make gcc-c++ && \
    dnf clean all

# Set environment variables
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /go/src/github.com/openshift/compliance-operator

ENV GOFLAGS=-mod=vendor

COPY . .
RUN make manager

# Step two: Containerize compliance-operator
FROM registry.access.redhat.com/ubi9/ubi-micro:latest

ENV OPERATOR=/usr/local/bin/compliance-operator \
    USER_UID=1001 \
    USER_NAME=compliance-operator

# Install operator binary
COPY --from=builder /go/src/github.com/openshift/compliance-operator/build/_output/bin/compliance-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}

