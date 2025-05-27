package mta

import (
	"io"
	"net/http"
	"time"

	pb "github.com/GlennTatum/prometheus-gtfs-exporter/mta/protobuf"
	"google.golang.org/protobuf/proto"
)

const (
	ACE = "https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-ace"
)

type NoFeedError string

func (n NoFeedError) Error() string {
	return "Feed not populated"
}

type client struct {
	c    http.Client
	feed *pb.FeedMessage
}

func NewClient() *client {
	return &client{
		c: http.Client{},
	}
}

/**
 * TODO
 * Returns the Departure Times of a current Stop
 */
type Departure struct {
}

func (client *client) StopDepartures(id string) []time.Time {
	var departures []time.Time = make([]time.Time, 0)

	for _, e := range client.feed.Entity {
		tripUpdate := e.TripUpdate
		if tripUpdate != nil {
			stu := tripUpdate.StopTimeUpdate
			for _, update := range stu {
				// NYCT_Train Extension
				// e_update := proto.GetExtension(update, pb.E_NyctStopTimeUpdate).(*pb.NyctStopTimeUpdate)
				if id == *update.StopId {
					departures = append(departures, time.Unix(*update.Departure.Time, 0))
				}
			}
		}
	}
	return departures
}

func (client *client) Get(s string) error {
	r, err := client.c.Get(s)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	msg := &pb.FeedMessage{}
	proto.Unmarshal(b, msg)
	client.feed = msg

	return nil
}
