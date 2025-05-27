package mta

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
)

func data(f string) string {
	return fmt.Sprintf("data/%s", f)
}

func StopsTXT() (map[string]string, error) {
	meta := make(map[string]string)
	data, err := os.ReadFile(data("stops.txt"))
	if err != nil {
		return meta, err
	}
	rd := bufio.NewReader(bytes.NewReader(data))
	stops := csv.NewReader(rd)
	records, err := stops.ReadAll()
	if err != nil {
		return meta, err
	}
	for i, r := range records {
		for range r {
			stop_id := records[i][0]
			stop_name := records[i][1]
			meta[stop_id] = stop_name
		}
	}
	return meta, nil
}
