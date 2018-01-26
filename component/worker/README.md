# how to use

```
// 启动dispatcher
worker.NewDispatcher(500).Run()

// 初始化job queue
worker.InitJobQueue(1000)

// 向job queue push job
worker.Push(job)
```