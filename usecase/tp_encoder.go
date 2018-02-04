package usecase

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
)

type TpEncoder interface {
	Encode(*domain.Counter, string) []string
	Header(string) string
}

var tpEncoder TpEncoder

func InitTpEncoder(e TpEncoder) {
	tpEncoder = e
}
