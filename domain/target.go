package domain

type Target struct {
	Url string
}

func NewTarget(url string) *Target {
	return &Target{
		Url: url,
	}
}
