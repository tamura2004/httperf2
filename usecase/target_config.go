package usecase

type TargetConfig struct {
	Url string
}

var Target TargetConfig

func InitTargetConfig(t TargetConfig) {
	Target = t
}
