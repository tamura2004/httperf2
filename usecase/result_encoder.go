package usecase

type ResultEncoder interface {
	Encode(*domain.Result, string) string
	Header() string
}

var resultEncoder ResultEncoder

func InitResultEncoder(e ResultEncoder) {
	resultEncoder = e
}
