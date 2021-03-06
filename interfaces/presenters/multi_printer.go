package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"log"
	"runtime"
	// "runtime/debug"
)

type MultiPrinter struct {
	Writer io.Writer
}

func (p *MultiPrinter) Print(c *domain.Counter) {
	fmt.Fprintf(p.Writer, "%s,%s,MULTI,%s,%d,%d\n", Date(), Time(), host.Name(), c.Multi, runtime.NumGoroutine())
	// debug.PrintStack()
}

func InitMultiPrinter() {
	p := &MultiPrinter{
		Writer: usecase.FileFactory.CreateTee("MULTI", "csv"),
	}
	fmt.Fprintln(p.Writer, "DATE,TIME,TYPE,HOSTNAME,MULTI,GOROUTINE")
	usecase.InitMultiPrinter(p)
	log.Println("init multi printer")

}
