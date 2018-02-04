package infra

import (
	"io"
	"log"
	"os"
)

func InitLogger() {
	file := CreateFile("log", "log")
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
}
