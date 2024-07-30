FROM 39z20094.c1.bhs5.container-registry.ovh.net/orca/firehose-core-koii@sha256:f023bccc67e4bcb1cf8f83d3c636e6c2a00108be0380a8001ae9fff5174a47cb as core

FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG VERSION="dev"
RUN apk --no-cache add git
RUN go build -v -ldflags "-X main.version=${VERSION}" ./cmd/firesol

# ADD firesol /app/firesol
#COPY tools/docker/motd_generic /etc/
#COPY tools/docker/motd_node_manager /etc/
#COPY tools/docker/99-firehose-solana.sh /etc/profile.d/
#COPY tools/docker/scripts/* /usr/local/bin
# On SSH connection, /root/.bashrc is invoked which invokes '/root/.bash_aliases' if existing,
# so we hijack the file to "execute" our specialized bash script
#RUN echo ". /etc/profile.d/99-firehose-solana.sh" > /root/.bash_aliases

FROM alpine:3

RUN apk --no-cache add \
        ca-certificates htop iotop sysstat \
        strace lsof curl jq tzdata

RUN mkdir -p /app/ && curl -Lo /app/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.12/grpc_health_probe-linux-amd64 && chmod +x /app/grpc_health_probe

WORKDIR /app

COPY --from=build /app/firesol /app/firesol

COPY --from=core /app/firecore /app/firecore

ENTRYPOINT ["/app/firesol"]
