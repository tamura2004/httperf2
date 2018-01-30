package mock

import (
	"io"
	"os"
)

type factory struct{}
type Writer struct {
	out *os.File
}

func NewFileFactory() *factory {
	return &factory{}
}

func (f *factory) CreateFile(pre, format, ext string) io.Writer {
	return &Writer{}
}

func (f *factory) CreateTeeFile(pre, format, ext string) io.Writer {
	return &Writer{}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (w *Writer) Close() error {
	return nil
}
