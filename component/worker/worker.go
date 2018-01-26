package worker

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	Quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}
}

func (worker Worker) Start() {
	go func() {
		for {
			worker.WorkerPool <- worker.JobChannel
			select {
			case job := <-worker.JobChannel:
				job.Proc(job.Data)
			case quit := <-worker.Quit:
				if quit {
					return
				}
			}
		}
	}()
}

func (worker Worker) stop() {
	go func() {
		worker.Quit <- true
	}()
}
