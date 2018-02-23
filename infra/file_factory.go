package infra

import (
	"fmt"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"log"
	"os"
	"time"
)

type FileFactory func(pre, ext string) io.Writer

func CreateFile(pre, ext string) io.Writer {
	timestamp := time.Now().Format("20060102_1504")
	name := fmt.Sprintf("%s_%s_%s.%s", timestamp, pre, Host.Name(), ext)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func (f FileFactory) Create(pre, ext string) io.Writer {
	return f(pre, ext)
}

func (f FileFactory) CreateTee(pre, ext string) io.Writer {
	file := f(pre, ext)
	return io.MultiWriter(file, os.Stdout)
}

func InitFileFactory() {
	usecase.FileFactory = FileFactory(CreateFile)
	log.Println("init file factory")

}
