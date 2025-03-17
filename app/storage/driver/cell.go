package driver

import "time"

type Cell struct {
	TTL     int       // 8
	AddDate time.Time // 24
	Value   string    // n\8  переписать на байты?
}

type Cell_O struct {
	Expired int64  // 8
	Value   string // n\8
}

type Cell_B struct {
	Expired int64  // 8
	Value   []byte // n\8
}
