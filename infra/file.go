package infra

import (
	"fmt"
	"log"
	"os"
	"time"
)

func CreateFile(pre, format, ext string) *os.File {
	timestamp := time.Now().Format(format)
	name := fmt.Sprintf("%s%s.%s", pre, timestamp, ext)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
