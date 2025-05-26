package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/GlennTatum/prometheus-gtfs-exporter/mta"
	pb "github.com/GlennTatum/prometheus-gtfs-exporter/mta/protobuf"

	"google.golang.org/protobuf/proto"
)

func main() {
	err := Exec()
	if err != nil {
		panic(err)
	}
}

func Exec() error {
	client := http.Client{}
	r, err := client.Get(mta.ACE)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var msg *pb.FeedMessage = &pb.FeedMessage{}
	proto.Unmarshal(b, msg)
	fmt.Println(msg)

	return nil
}
