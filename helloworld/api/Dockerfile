FROM golang:1.19 as build

WORKDIR /src
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY . /src
RUN ls -al
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app /src/api

FROM gcr.io/distroless/base:debug-nonroot
WORKDIR /

COPY --from=build /app /app
COPY --from=build /src/api/etc /etc

ENTRYPOINT ["/app"]