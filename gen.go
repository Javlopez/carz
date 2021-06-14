package carz


//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative carz.proto
/*
//go:generate protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative carz.proto
*/
