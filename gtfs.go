package gtfs

import (
	"net/http"
)

type GTFS struct {
	r *http.Client
}
