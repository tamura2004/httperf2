package presenters

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"github.com/tamura2004/httperf2/usecase"
	"log"
)

const AVERAGE_HEADER = "DATE,TIME,TYPE,HOSTNAME,DURATION,COUNT"

type AveragePrinter struct{}

func (p *AveragePrinter) Print(c *domain.Counter) {
	file := usecase.FileFactory.CreateTee("AVERAGE", "csv")
	fmt.Fprintln(file, AVERAGE_HEADER)
	fmt.Fprintf(file, "%s,%s,AVERAGE,%s,%s,%d", Date(), Time(), host.Name(), c.Average(), c.Count)
}

func InitAveragePrinter() {
	usecase.InitAveragePrinter(&AveragePrinter{})
	log.Println("init average printer")

}
