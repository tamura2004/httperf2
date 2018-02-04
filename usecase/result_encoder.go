package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type resultEncoder interface {
	Encode(*domain.Result) string
	Header() string
}

var ResultEncoder resultEncoder

func InitResultEncoder(e resultEncoder) {
	ResultEncoder = e
}
