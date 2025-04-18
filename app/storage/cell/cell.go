package cell

import "time"

type Cell struct {
	expired int64 // 8
	Value   []byte
}

func NewCell(v []byte, ttl int) *Cell {
	c := &Cell{
		Value: make([]byte, len(v)),
	}

	c.Value = v

	if ttl > 0 {
		c.expired = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	}

	return c
}

func (c *Cell) IsExpired() bool {
	return c.expired > 0 && c.expired < time.Now().Unix()
}

// Refill перезаполняет ячейку
func (c *Cell) Refill(v []byte, ttl int) {
	c.Value = v

	if ttl > 0 {
		c.expired = time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	} else {
		c.expired = 0
	}
}
