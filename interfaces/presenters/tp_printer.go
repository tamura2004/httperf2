package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"log"
)

type TpPrinter struct {
	TpEncoder
}

func (p *TpPrinter) Print(c *domain.Counter, name string) {
	file := usecase.FileFactory.CreateTee(name, "csv")
	fmt.Fprintln(file, p.Header(name))
	for _, line := range p.Encode(c, name) {
		fmt.Fprintln(file, line)
	}
}

func InitTpPrinter() {
	p := &TpPrinter{
		TpEncoder: TpEncoder{},
	}
	usecase.InitTpPrinter(p)
	log.Println("init tp printer")
}
