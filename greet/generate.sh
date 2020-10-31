#! /bin/zsh

protoc greetpb/greet.proto --go_out=plugins=grpc:.