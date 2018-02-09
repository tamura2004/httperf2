package presenters

import (
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

type HttpBodyPrinter struct {
	Once sync.Once
}

func (p *HttpBodyPrinter) Print(r io.ReadCloser) {
	defer r.Close()

	// if p.OncePrinted {
	// 	return
	// }

	p.Once.Do(func() {
		file := usecase.FileFactory.CreateTee("HTTP_BODY", "html")
		io.Copy(file, r)
	})
	io.Copy(ioutil.Discard, r)

	// p.OncePrinted = true
}

func InitHttpBodyPrinter() {
	usecase.InitHttpBodyPrinter(&HttpBodyPrinter{sync.Once{}})
	log.Println("init http body printer")
}
