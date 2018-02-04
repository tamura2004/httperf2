package usecase

import (
	"github.com/tamura2004/httperf2/domain"
)

type ScinarioConfig struct {
	WorkerNum int             // ワーカー数
	Count     int             // 各ワーカのジョブ実行回数
	RampUp    domain.Duration // ワーカーの立ち上げ間隔
	Interval  domain.Duration // 各ワーカーのジョブ実施間隔
}

var Scinario ScinarioConfig

func InitScinarioConfig(s ScinarioConfig) {
	Scinario = s
}
