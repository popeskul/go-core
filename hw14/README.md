# Install
`go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

# Lunch
PB: `protoc -I . --go_out=./ ./messages_proto/messages.proto`

gRPC: `protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative messages_proto/messages.proto`

or

`make protoc`
