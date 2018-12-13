package payloadcontrol

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	jobQueue   chan Job
	workerPool chan chan Job
	maxWorkers int
}

func (d *Dispatcher) Push(payload Payload) {
	job := Job{Payload: payload}
	d.jobQueue <- job
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.workerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
