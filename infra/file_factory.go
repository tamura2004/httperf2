package infra

import (
	"fmt"
	"github.com/tamura2004/httperf2/usecase"
	"io"
	"log"
	"os"
	"time"
)

func CreateFile(pre, ext string) *os.File {
	timestamp := time.Now().Format("2006_01_02_15_04_05_")
	name := fmt.Sprintf("%s%s.%s", timestamp, pre, ext)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type fileFactory struct{}

func InitFileFactory() {
	usecase.InitFileFactory(&fileFactory{})
	log.Println("init file factory")

}

func (f *fileFactory) Create(pre, ext string) io.Writer {
	return CreateFile(pre, ext)
}

func (f *fileFactory) CreateTee(pre, ext string) io.Writer {
	file := CreateFile(pre, ext)
	return io.MultiWriter(file, os.Stdout)
}
