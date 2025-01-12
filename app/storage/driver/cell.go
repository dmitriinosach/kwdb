package driver

import "time"

type Cell struct {
	Value   string
	Key     string
	TTL     int
	AddDate time.Time
}
