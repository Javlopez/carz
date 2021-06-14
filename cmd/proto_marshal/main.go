package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Javlopez/carz/pb"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	loc := pb.Location{
		Lat: -22.95,
		Lng: -43.21,
	}

	data, err := proto.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}

	//os.Stdout.Write(data)

	fmt.Println("protobuf size:", len(data))
	jdata, err :=  json.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON size", len(jdata))

	var loc2 pb.Location

	if err := proto.Unmarshal(data, &loc2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", &loc2)

	s := pb.Status_FREE
	fmt.Printf("free: %s (%d)\n", s, s)

	car := &pb.Car{
		Id:       "95",
		Driver:   "Lightning McQueen",
		Location: &loc,
		Status:   pb.Status_FREE,
		Updated: timestamppb.Now(),
	}

	fmt.Println(car)

	t :=  car.Updated.AsTime()
	fmt.Println(t)

	tz, err := time.LoadLocation("America/Mexico_City")
	mx := t.In(tz)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mx)

	data, err = proto.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("proto size %d\n", len(data))

	jdata, err = json.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON size %d\n", len(jdata))
	fmt.Println(string(jdata))

	jdata, err = protojson.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jdata))

	///fmt.Printf("%#v\n", car)
	//fmt.Printf("%#v\n", car) not that readable

}