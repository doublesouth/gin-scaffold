package worker

type Job struct {
	Data interface{}
	Proc func(interface{})
}

var jobQueue chan Job

// 初始化 job queue
func InitJobQueue(max int) {
	jobQueue = make(chan Job, max)
}

func Push(job Job) {
	jobQueue <- job
}
