## protobuf score

protoc --go_out=proto/score proto/score/score.proto

## grpc greeter
go get -u github.com/golang/protobuf/protoc-gen-go

protoc --go_out=plugins=grpc:proto/greeter proto/greeter/greeter.proto

## grpc hello

go get google.golang.org/protobuf/cmd/protoc-gen-go

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloworld/helloworld.proto

## grpc routeguide

go get google.golang.org/protobuf/cmd/protoc-gen-go

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/routeguide/route_guide.proto
