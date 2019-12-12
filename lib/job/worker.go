package job

import (
	"fmt"
)

type Job interface {
	Payload() error
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.

			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Payload(); err != nil {
					//					log.Errorf("Error uploading to S3: %s", err.Error())
					fmt.Print(err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop

				return

			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true

	}()
}
