package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type Collector struct {
	Counter *domain.Counter
	Result  chan *domain.Result
	*Reporter
}

func NewCollector(res chan *domain.Result, rep *Reporter) *Collector {
	return &Collector{
		Counter:  domain.NewCounter(),
		Result:   res,
		Reporter: rep,
	}
}

func (c *Collector) Run() {
	for result := range c.Result {
		c.CollectResult(result)
	}
	c.Reporter.ReportTPM(c.Counter)
	c.Reporter.ReportTPS(c.Counter)
}

func (c *Collector) CollectResult(r *domain.Result) {
	switch r.Check {
	case domain.START:
		c.Counter.IncMulti()
		c.ReportResult(r)

	case domain.CONNECT:
		c.ReportResult(r)

	case domain.FINISH:
		c.ReportResult(r)
		c.Counter.IncTp()
		c.Counter.DecMulti()
	}
}
