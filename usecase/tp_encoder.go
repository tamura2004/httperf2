package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type tpEncoder interface {
	Encode(*domain.Counter, string) []string
	Header(string) string
}

var TpEncoder tpEncoder

func InitTpEncoder(e tpEncoder) {
	TpEncoder = e
}
