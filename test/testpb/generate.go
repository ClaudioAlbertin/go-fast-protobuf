//go:generate protoc --go_out=. --plugin=../../cmd/protoc-gen-lazy/protoc-gen-lazy --lazy_out=. --go-vtproto_out=. --go-vtproto_opt=features=marshal+unmarshal+size test.proto
package testpb
