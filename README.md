# prometheus-gtfs-exporter
A pluggable General Transit Feed Specification Prometheus Exporter


### Setup

1. Install the protoc package
```bash
apt install -y protobuf-compiler
```
2. Install the go protobuf plugin
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
| The binary will be installed to `$GOPATH/bin`
