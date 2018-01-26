package starter

import (
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
	"github.com/doublesouth/gin-scaffold/component/worker"
)

func TestWorker2(t *testing.T) {
	worker.InitJobQueue(10000)
	worker.NewDispatcher(5000).Run()
	log.Print("worker started")

	jobSize := 30000
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(jobSize)
	for i := 0; i < jobSize; i++ {
		job := worker.Job{
			Data: "data num " + strconv.Itoa(i),
			Proc: func(data interface{}) {
				log.Printf("proc data: %s \n", data)

				// TODO 模拟http请求，测试当有大量http请求时，服务器各资源占用率
				time.Sleep(2 * time.Second)
				waitGroup.Done()
			},
		}
		worker.Push(job)
	}

	waitGroup.Wait()
	log.Print("all done")
}
