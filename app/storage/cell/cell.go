package cell

import "time"

type Cell struct {
	expired int64  // 8
	Value   string // n\8
}

type Cell_O struct {
	expired int64  // 8
	Value   []byte // n\8
}

func NewCell_o(v []byte, ttl int) *Cell_O {
	c := &Cell_O{
		Value: make([]byte, len(v)),
	}

	if ttl > 0 {
		c.expired = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	return c
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
