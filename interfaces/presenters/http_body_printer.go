package presenters

import (
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"io/ioutil"
	"log"
)

type HttpBodyPrinter struct {
	OncePrinted bool
}

func (p *HttpBodyPrinter) Print(r io.ReadCloser) {
	defer r.Close()

	if p.OncePrinted {
		io.Copy(ioutil.Discard, r)
		return
	}

	file := usecase.FileFactory.CreateTee("HTTP_BODY", "html")
	io.Copy(file, r)
	p.OncePrinted = true
}

func InitHttpBodyPrinter() {
	usecase.InitHttpBodyPrinter(&HttpBodyPrinter{false})
	log.Println("init http body printer")
}
