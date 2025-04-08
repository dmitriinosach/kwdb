package cell

import "time"

type Cell struct {
	expired int64  // 8
	Value   string // n\8
}

func NewCell(v string, ttl int) *Cell {
	c := &Cell{
		Value: v,
	}

	if ttl > 0 {
		c.expired = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	return c
}

func (c *Cell) IsExpired() bool {
	return c.expired > 0 && c.expired < time.Now().Unix()
}

// Refill перезаполняет ячейку
func (c *Cell) Refill(v string, ttl int) {
	c.Value = v

	if ttl > 0 {
		c.expired = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	} else {
		c.expired = 0
	}
}
