package presenters

import (
	"github.com/tamura2004/httperf2/domain"
)

func Date() string {
	return domain.Time.Now().Format("2006-01-02")
}

func Time() string {
	return domain.Time.Now().Format("15:04:05")
}
