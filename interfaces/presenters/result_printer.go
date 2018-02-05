package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"log"
)

type ResultPrinter struct {
	ResultEncoder
	Writer io.Writer
}

func (p *ResultPrinter) Print(r *domain.Result) {
	fmt.Fprintln(p.Writer, p.Encode(r))
}

func InitResultPrinter() {
	p := &ResultPrinter{
		ResultEncoder: ResultEncoder{},
		Writer:        usecase.FileFactory.CreateTee("RESULT", "csv"),
	}
	fmt.Fprintln(p.Writer, p.Header())
	usecase.InitResultPrinter(p)
	log.Println("init result printer")
}
