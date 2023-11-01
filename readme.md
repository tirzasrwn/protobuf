# protobuf

## About

A note about protocol buffer implementation in go.

## Requirement

- protoc
- protoc-gen-go

```sh
sudo apt install protobuf-compiler
protoc --version
sudo apt install protoc-gen-go
protoc-gen-go --version
```

## Run

```sh
# to generate *.pb.go
make proto
# to run
make
```
