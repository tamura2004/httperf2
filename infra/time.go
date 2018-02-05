package infra

import (
	"github.com/tamura2004/httperf2/domain"
	"log"
	"time"
)

type Time struct{}

func InitTime() {
	domain.InitTime(&Time{})
	log.Println("init time")

}

func (*Time) Now() time.Time {
	return time.Now()
}

func (*Time) Since(t time.Time) time.Duration {
	return time.Since(t)
}
