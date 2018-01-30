package domain

import (
	"math/rand"
	"time"
)

var Now = time.Now

func init() {
	rand.Seed(Now().UnixNano())
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
	time.Sleep(d.ExpRand())
}

func (d *Duration) ExpRand() time.Duration {
	return time.Duration(float64(d.Duration) * rand.ExpFloat64())
}
