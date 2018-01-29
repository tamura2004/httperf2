package domain

type CheckPoint int

const (
	START CheckPoint = iota + 1
	CONNECT
	FINISH
)

func (c CheckPoint) String() string {
	switch c {
	case START:
		return "START"
	case CONNECT:
		return "CONNECT"
	case FINISH:
		return "FINISH"
	default:
		return "undefined"
	}
}
