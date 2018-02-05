package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type resultprinter interface {
	Print(*domain.Result)
}

var ResultPrinter resultprinter

func InitResultPrinter(p resultprinter) {
	ResultPrinter = p
}
