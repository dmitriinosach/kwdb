package storage

import "time"

type Cell struct {
	Value     string
	Key       string
	TTL       int
	AddDate   time.Time
	IsLocked  bool
	IsExpired bool
}