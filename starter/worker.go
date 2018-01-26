package starter

import (
	"github.com/golang/glog"
	"github.com/doublesouth/gin-scaffold/component/worker"
	"github.com/doublesouth/gin-scaffold/conf"
)

// 启动worker，向job queue取job执行
func Worker() {
	worker.InitJobQueue(conf.Conf.Worker.JobSize)
	worker.NewDispatcher(conf.Conf.Worker.WorkerSize).Run()
	glog.Info("worker started")
}
