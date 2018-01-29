package domain

import (
	"time"
)

type Unit int

const (
	TPS Unit = iota + 1
	TPM
)

type TimeStamp string

type Counter struct {
	Tr    []map[TimeStamp]int
	Count int
	Multi int
	Total time.Duration
}

func (u Unit) String() string {
	switch u {
	case TPS:
		return "TPS"
	case TPM:
		return "TPM"
	default:
		return "undefined"
	}
}
