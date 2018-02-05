package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"log"
)

type AveragePrinter struct{}

func (p *AveragePrinter) Header() string {
	return "DATE,TIME,TYPE,DURATION,COUNT"
}

func (p *AveragePrinter) Print(c *domain.Counter) {
	file := usecase.FileFactory.CreateTee("AVERAGE", "csv")
	fmt.Fprintf(file, "%s,%s,AVERAGE,%s,%d", Date(), Time(), c.Average(), c.Count)
}

func InitAveragePrinter() {
	usecase.InitAveragePrinter(&AveragePrinter{})
	log.Println("init average printer")

}
