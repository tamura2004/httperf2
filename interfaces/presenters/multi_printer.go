package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"log"
)

type MultiPrinter struct {
	Writer io.Writer
}

func (p *MultiPrinter) Print(c *domain.Counter) {
	fmt.Fprintf(p.Writer, "%s,%s,MULTI,%d\n", Date(), Time(), c.Multi)
}

func InitMultiPrinter() {
	p := &MultiPrinter{
		Writer: usecase.FileFactory.CreateTee("MULTI", "csv"),
	}
	fmt.Fprintln(p.Writer, "DATE,TIME,TYPE,MULTI")
	usecase.InitMultiPrinter(p)
	log.Println("init multi printer")

}
