package usecase

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"io"
)

type Reporter struct {
	Logfile io.Writer
}

func NewReporter() *Reporter {
	logfile := FileFactory.CreateTee("log", "csv")
	fmt.Fprintln(logfile, ResultEncoder.Header())
	return &Reporter{
		Logfile: logfile,
	}
}

func (re *Reporter) ReportResult(r *domain.Result) {
	fmt.Fprintf(re.Logfile, "%s\n", ResultEncoder.Encode(r))
}

func (re *Reporter) ReportTp(c *domain.Counter, name string) {
	file := FileFactory.CreateTee(name, "csv")
	fmt.Fprintln(file, TpEncoder.Header(name))
	for _, row := range TpEncoder.Encode(c, name) {
		fmt.Fprintln(file, row)
	}
}

func (re *Reporter) ReportTPM(c *domain.Counter) {
	re.ReportTp(c, "TPM")
}

func (re *Reporter) ReportTPS(c *domain.Counter) {
	re.ReportTp(c, "TPS")
}

func (re *Reporter) ReportAverage(c *domain.Counter) {
	AveragePrinter.Print(c)
}
