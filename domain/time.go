package domain

import (
	"math/rand"
	"time"
)

type clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

var _time clock

func InitTime(c clock) {
	_time = c
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
	time.Sleep(d.ExpRand().Duration)
}

func (d *Duration) ExpRand() *Duration {
	return &Duration{
		time.Duration(float64(d.Duration) * rand.ExpFloat64()),
	}
}

func (d *Duration) Times(n int) *Duration {
	return &Duration{
		d.Duration * time.Duration(n),
	}
}
