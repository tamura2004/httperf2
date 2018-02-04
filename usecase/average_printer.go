package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type averagePrinter interface {
	Print(*domain.Counter)
}

var AveragePrinter averagePrinter

func InitAveragePrinter(p averagePrinter) {
	AveragePrinter = p
}
