package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Worker struct {
	ID       int
	JobChan  chan Job
	QuitChan chan bool
}

func NewWorker(id int, jobChan chan Job) *Worker {
	return &Worker{
		ID:       id,
		JobChan:  jobChan,
		QuitChan: make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobChan:
				fmt.Printf("Worker %d processing job %d: %s\n", w.ID, job.ID, job.Data)
				time.Sleep(500 * time.Millisecond) // Simulate work
				fmt.Printf("Worker %d completed job %d\n", w.ID, job.ID)
			case <-w.QuitChan:
				fmt.Printf("Worker %d quitting\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.QuitChan <- true
}

func main() {
	jobChan := make(chan Job, 10)

	// Start workers
	workers := make([]*Worker, 3)
	for i := 0; i < 3; i++ {
		workers[i] = NewWorker(i+1, jobChan)
		workers[i].Start()
	}

	// Send jobs
	for i := 1; i <= 5; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("Job data %d", i),
		}
		jobChan <- job
	}

	// Wait a bit
	time.Sleep(3 * time.Second)

	// Stop workers
	for _, worker := range workers {
		worker.Stop()
	}
}
