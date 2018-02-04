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
	logfile := fileFactory.CreateTee("log", "csv")
	fmt.Fprintln(logfile, resultEncoder.Header())
	return &Reporter{
		Logfile: logfile,
	}
}

func (re *Reporter) ReportResult(r *domain.Result) {
	fmt.Fprintf(re.Logfile, "%s\n", resultEncoder.Encode(r))
}

func (re *Reporter) ReportTp(c *domain.Counter, name string) {
	file := fileFactory.CreateTee(name, "csv")
	fmt.Fprintln(file, tpEncoder.Header(name))
	for _, row := range tpEncoder.Encode(c, name) {
		fmt.Fprintln(file, row)
	}
}

func (re *Reporter) ReportTPM(c *domain.Counter) {
	re.ReportTp(c, "TPM")
}

func (re *Reporter) ReportTPS(c *domain.Counter) {
	re.ReportTp(c, "TPS")
}
