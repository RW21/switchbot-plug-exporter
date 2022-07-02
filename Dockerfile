FROM golang:1.18 as builder

ARG GOARCH
ARG GOOS

COPY . /src
WORKDIR /src
RUN GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o exporter .

FROM scratch
COPY --from=builder /src/exporter /exporter
EXPOSE 9101
ENTRYPOINT ["/exporter"]
