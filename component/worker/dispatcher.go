package worker

type Dispatcher struct {
	MaxWorker  int
	WorkerPool chan chan Job
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func NewDispatcher(maxWorker int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorker)
	return &Dispatcher{
		WorkerPool: workerPool,
		MaxWorker:  maxWorker,
	}
}
