package domain

import (
	"time"
)

type Counter struct {
	TPS   map[string]int
	TPM   map[string]int
	Count int
	Multi int
	Total time.Duration
}

func NewCounter() *Counter {
	return &Counter{
		TPS: make(map[string]int),
		TPM: make(map[string]int),
	}
}

func (c *Counter) IncTp() {
	c.IncTPM()
	c.IncTPS()
}

func (c *Counter) IncTPM() {
	t := time.Now().Format("2006-01-02,15:04")
	c.TPM[t]++
}

func (c *Counter) IncTPS() {
	t := time.Now().Format("2006-01-02,15:04:05")
	c.TPS[t]++
}

func (c *Counter) IncMulti() {
	c.Multi++
}

func (c *Counter) DecMulti() {
	c.Multi--
}
