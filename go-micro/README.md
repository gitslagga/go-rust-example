## protobuf score

protoc --go_out=proto/score proto/score/score.proto

## grpc greeter

protoc --go_out=plugins=grpc:proto/greeter proto/greeter/greeter.proto

