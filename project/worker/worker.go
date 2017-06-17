package main

import (
	"log"
	"time"
)

const (
	MAX_QUEUE   = 1000
	MAX_WORKERS = 100
)

type Job struct {
	Id int
}

func (p *Job) Working() error {
	//todo...
	log.Println("My working id:", p.Id)

	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			return nil
		}
	}
	return nil
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MAX_QUEUE)
}

// Worker represents the worker that executes the job
type Worker struct {
	Working    bool
	JobChannel chan Job
	quit       chan bool
}

func NewWorker() Worker {
	return Worker{
		Working:    false,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobChannel:
				log.Println("start work :", job.Id)

				// we have received a work request.
				if err := job.Working(); err != nil {
					log.Println("Error uploading to S3: %s", err.Error())
				}
				w.Working = false
			case <-w.quit:
				w.Working = false
				// we have received a signal to stop
				return
			}
		}
	}()
}

func (w *Worker) SetWorking() {
	w.Working = true
}

// Stop signals the worker to stop listening for work requests.
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
