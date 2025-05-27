package main

import (
	"fmt"

	"github.com/GlennTatum/prometheus-gtfs-exporter/mta"
)

func main() {
	err := Exec()
	if err != nil {
		panic(err)
	}
}

func Exec() error {
	// client := mta.NewClient()
	// err := client.Get(mta.ACE)
	// if err != nil {
	// 	return err
	// }
	// d := client.StopDepartures("A02S")
	// for _, v := range d {
	// 	fmt.Println(v)
	// }
	fmt.Println(mta.StopsTXT())

	return nil
}
