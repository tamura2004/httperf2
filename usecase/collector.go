package usecase

import (
	"github.com/tamura2004/httperf2/domain"
	"log"
	"time"
)

type Collector struct {
	Counter *domain.Counter
	Timer   *time.Timer
}

func NewCollector() *Collector {
	return &Collector{
		Counter: domain.NewCounter(),
		Timer:   time.NewTimer(Scinario.Timeout.Duration),
	}
}

func (c *Collector) Run() {
	CheckInit()

	go c.CollectMulti()

	go c.WatchTimeOut()

	for result := range resultChan {
		c.CollectResult(result)
	}
	TpPrinter.Print(c.Counter, "TPM")
	TpPrinter.Print(c.Counter, "TPS")
	AveragePrinter.Print(c.Counter)
}

func (c *Collector) CollectMulti() {
	for range time.Tick(10 * time.Second) {
		MultiPrinter.Print(c.Counter)
	}
}

func (c *Collector) WatchTimeOut() {
	for range c.Timer.C {
		close(resultChan)
		log.Println("time out")
	}
}

func (c *Collector) CollectResult(r *domain.Result) {
	CheckInit()

	switch r.Check {
	case domain.START:
		c.Counter.IncMulti()
		ResultPrinter.Print(r)

	case domain.CONNECT:
		ResultPrinter.Print(r)

	case domain.FINISH:
		c.Timer.Reset(Scinario.Timeout.Duration)
		ResultPrinter.Print(r)
		c.Counter.IncTp()
		c.Counter.DecMulti()

		c.Counter.AddDuration(r.Duration)
	}
}
