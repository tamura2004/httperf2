package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

var resultChan chan *domain.Result

func InitResultChan() {
	resultChan = make(chan *domain.Result, 20000)
}
