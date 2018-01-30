package domain

import (
	"time"
)

var now = time.Now

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
	time.Sleep(d.Duration)
}
