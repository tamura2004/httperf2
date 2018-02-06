package usecase

import (
	"io"
)

type httpBodyPrinter interface {
	Print(io.ReadCloser)
}

var HttpBodyPrinter httpBodyPrinter

func InitHttpBodyPrinter(p httpBodyPrinter) {
	HttpBodyPrinter = p
}
