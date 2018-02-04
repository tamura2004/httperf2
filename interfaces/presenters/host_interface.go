package presenters

type Host interface {
	Name() string
}

var host Host

func InitHost(h Host) {
	host = h
}
