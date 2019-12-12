package job

import (
	"fmt"
)

type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
	maxQueue   int
}

func NewDispatcher(maxWorkers int, maxQueue int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	JobQueue = make(chan Job, maxQueue)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers, maxQueue: maxQueue}
}

func (d *Dispatcher) Run() {
	// starting n number of workers

	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			fmt.Println(job)

			go func(job Job) {
				// try to obtain a worker job channel that is available.

				// this will block until a worker is idle

				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel

				jobChannel <- job
			}(job)
		}
	}
}
