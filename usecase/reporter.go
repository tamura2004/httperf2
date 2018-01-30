package usecase

import (
	"fmt"
	"github.com/tamura2004/httperf2/domain"
	"io"
)

type ResultEncoder interface {
	EncodeResult(*domain.Result, string) string
	HeaderResult() string
}

type TpEncoder interface {
	EncodeTp(*domain.Counter, string, string) []string
	HeaderTp(string) string
}

type Reporter struct {
	Logfile io.Writer
	FileFactory
	ResultEncoder
	TpEncoder
	Hostname string
}

type FileFactory interface {
	CreateFile(pre, format, ext string) io.Writer
	CreateTeeFile(pre, format, ext string) io.Writer
}

func NewReporter(f FileFactory, e ResultEncoder, t TpEncoder, h string) *Reporter {
	logfile := f.CreateTeeFile("log", "20060102150405", "csv")
	fmt.Fprintln(logfile, e.HeaderResult())
	return &Reporter{
		Logfile:       logfile,
		FileFactory:   f,
		ResultEncoder: e,
		TpEncoder:     t,
		Hostname:      h,
	}
}

func (re *Reporter) ReportResult(r *domain.Result) {
	fmt.Fprintf(re.Logfile, "%s\n", re.EncodeResult(r, re.Hostname))
}

func (re *Reporter) ReportTp(c *domain.Counter, name string) {
	file := re.CreateTeeFile(name, "20060102150405", "csv")
	fmt.Fprintln(file, re.HeaderTp(name))
	for _, row := range re.EncodeTp(c, name, re.Hostname) {
		fmt.Fprintln(file, row)
	}
}

func (re *Reporter) ReportTPM(c *domain.Counter) {
	re.ReportTp(c, "TPM")
}

func (re *Reporter) ReportTPS(c *domain.Counter) {
	re.ReportTp(c, "TPS")
}
