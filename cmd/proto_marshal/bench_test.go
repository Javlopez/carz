package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Javlopez/carz/pb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	car = &pb.Car{
		Id:       "95",
		Driver:   "Lightning McQueen",
		Location: &pb.Location{
			Lat: -22.95,
			Lng: -43.21,
		},
		Status:   pb.Status_FREE,
		Updated: timestamppb.Now(),
	}
)


func BenchmarkJSON(b *testing.B) {
	for i := 0; i< b.N; i++{
		_, err :=  json.Marshal(car)
		if err != nil {
			log.Fatal(err)
		}
	}
}


func BenchmarkProto(b *testing.B) {
	for i := 0; i< b.N; i++{
		_, err :=  proto.Marshal(car)
		if err != nil {
			log.Fatal(err)
		}
	}
}
