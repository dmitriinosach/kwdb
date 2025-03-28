package driver

import "time"

type Cell struct {
	expired int64  // 8
	Value   string // n\8
}

func NewCell(v string, ttl int) *Cell {
	return &Cell{
		expired: time.Now().Add(time.Duration(ttl) * time.Second).Unix(),
		Value:   v,
	}
}

func (c *Cell) IsExpired() bool {
	return time.Now().Unix() > c.expired
}
