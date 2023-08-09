FROM golang:1.21 as build

WORKDIR /src
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY . /src
RUN ls -al
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app /src/rpc

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.14 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64
RUN chmod +x /bin/grpc_health_probe

FROM gcr.io/distroless/base:debug-nonroot
WORKDIR /

COPY --from=build /app /app
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe

ENTRYPOINT ["/app"]