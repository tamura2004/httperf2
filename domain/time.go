package domain

import (
	"math/rand"
	"time"
)

var now = time.Now

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (d *Duration) String() string {
	return d.Duration.String()
}

func (d *Duration) Sleep() {
	time.Sleep(d.Duration * time.Duration(rand.ExpFloat64()))
}

// time.Sleep(d * time.Duration(rand.ExpFloat64()))
