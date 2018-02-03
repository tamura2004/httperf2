package infra

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func CreateFile(pre, format, ext string) *os.File {
	timestamp := time.Now().Format(format)
	name := fmt.Sprintf("%s%s.%s", timestamp, pre, ext)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type FileFactory struct{}

func NewFileFactory() *FileFactory {
	return &FileFactory{}
}

func (f *FileFactory) CreateFile(pre, format, ext string) io.Writer {
	return CreateFile(pre, format, ext)
}

func (f *FileFactory) CreateTeeFile(pre, format, ext string) io.Writer {
	file := CreateFile(pre, format, ext)
	return io.MultiWriter(file, os.Stdout)
}
