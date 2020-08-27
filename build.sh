protoc --proto_path ../../../ -I=./proto --go_out=plugins=grpc:./proto proto/recordscores.proto
mv proto/github.com/brotherlogic/recordscores/proto/* ./proto
