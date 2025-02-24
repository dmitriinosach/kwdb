package driver

import "time"

type Cell struct {
	Value   string
	TTL     int
	AddDate time.Time
}
