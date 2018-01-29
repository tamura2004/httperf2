package domain

type Scinario struct {
	Worker   int      // ワーカー数
	Count    int      // 各ワーカのジョブ実行回数
	RampUp   Duration // ワーカーの立ち上げ間隔
	Interval Duration // 各ワーカーのジョブ実施間隔
}
