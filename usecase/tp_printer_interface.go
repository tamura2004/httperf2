package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type tpPrinter interface {
	Print(*domain.Counter, string)
}

var TpPrinter tpPrinter

func InitTpPrinter(p tpPrinter) {
	TpPrinter = p
}
