package infrastructure

import (
	"io"
	"time"
)

type SlowReader struct {
	delay time.Duration
	r     io.ReadCloser
}

func (sr SlowReader) Read(p []byte) (int, error) {
	time.Sleep(sr.delay)
	return sr.r.Read(p[:1])
}

func (sr SlowReader) Close() error {
	return sr.r.Close()
}

func NewSlowReader(r io.ReadCloser, bps int) io.ReadCloser {
	delay := time.Second / time.Duration(bps)
	return SlowReader{
		r:     r,
		delay: delay,
	}
}
