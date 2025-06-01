package main

import (
	"github.com/GlennTatum/prometheus-gtfs-exporter/mta"
	"github.com/GlennTatum/prometheus-gtfs-exporter/mta/exporter"
)

func main() {
	err := Exec()
	if err != nil {
		panic(err)
	}
}

func testing() {
	client := mta.NewClient()
	client.Get(mta.ACE)
	// d := client.StopDepartures("125S")
	// for _, stu := range d {
	// 	fmt.Println(stu.Departure_Time.Unix())
	// }
	// fmt.Println(mta.StopsTXT())

	exporter.Collect_departure_frequency(client, "A02S")
}

func Exec() error {

	// testing()
	exporter.Exec()
	return nil
}
