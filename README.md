# adder-service
A simple TWIRP adder service.

### Generate the Twirp Service
1. Install [protobuf](https://github.com/golang/protobuf#installation)
2. Install [twirp](https://twitchtv.github.io/twirp/docs/install.html#with-go-get)
3. Run protoc
```shell script
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/adder/service.proto
```

### Run the Service Locally
```shell
go run ./cmd/adder/main.go
```
