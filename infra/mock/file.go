package mock

import (
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"os"
)

type factory struct{}
type Writer struct {
	out *os.File
}

func InitFileFactory() {
	usecase.InitFileFactory(&factory{})
}

func (f *factory) Create(pre, ext string) io.Writer {
	return &Writer{}
}

func (f *factory) CreateTee(pre, ext string) io.Writer {
	return &Writer{}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (w *Writer) Close() error {
	return nil
}
