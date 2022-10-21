FROM golang:1.19 as build

WORKDIR /src
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY helloworld/api /src
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app /src

FROM gcr.io/distroless/base:debug-nonroot

COPY --from=build /app /app
COPY --from=build /src/etc /app/etc

ENTRYPOINT ["/app"]