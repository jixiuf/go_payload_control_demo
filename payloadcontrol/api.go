package payloadcontrol

type Payload interface {
	Handle() error
}

type Logger interface {
	Errorf(format string, v ...interface{})
	Error(v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
}

func NewDispatcher(maxWorkers, maxQueue int) *Dispatcher {
	return &Dispatcher{
		maxWorkers: maxWorkers,
		workerPool: make(chan chan Job, maxWorkers),
		jobQueue:   make(chan Job, maxQueue),
	}
}

var log Logger

func RegisterLogger(l Logger) {
	log = l
}
