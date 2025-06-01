package mta

import (
	"io"
	"net/http"
	"time"

	pb "github.com/GlennTatum/prometheus-gtfs-exporter/mta/protobuf"
	"google.golang.org/protobuf/proto"
)

type StationRouteLabel string

const (
	ACE StationRouteLabel = "https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-ace"
)

type NoFeedError string

func (n NoFeedError) Error() string {
	return "Feed not populated"
}

type Client struct {
	c    http.Client
	feed *pb.FeedMessage
}

func NewClient() *Client {
	return &Client{
		c: http.Client{},
	}
}

type Departure struct {
	Departure_Time time.Time
}

func (client *Client) StopDepartures(id string) []Departure {
	var departures []Departure = make([]Departure, 0)

	for _, e := range client.feed.Entity {
		tripUpdate := e.TripUpdate
		if tripUpdate != nil {
			stu := tripUpdate.StopTimeUpdate
			for _, update := range stu {
				// NYCT_Train Extension
				// e_update := proto.GetExtension(update, pb.E_NyctStopTimeUpdate).(*pb.NyctStopTimeUpdate)
				if id == *update.StopId {
					d := Departure{
						Departure_Time: time.Unix(*update.Departure.Time, 0),
						// departure_delay: *update.Departure.Delay,
					}
					departures = append(departures, d)
				}
			}
		}
	}
	return departures
}

func (client *Client) Get(s StationRouteLabel) error {
	r, err := client.c.Get(string(s))
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
