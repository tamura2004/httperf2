package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type multiPrinter interface {
	Print(*domain.Counter)
}

var MultiPrinter multiPrinter

func InitMultiPrinter(p multiPrinter) {
	MultiPrinter = p
}
