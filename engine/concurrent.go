package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	// 把in放到Scheduler里面
	ConfigureMasterWorkerChan(chan Request)
	// workerChan：从外界告诉我们有一个worker它可以负责去接收它可以负责去接收request
	WorkerReady(chan Request)
	// 启动一个总控的goroutine
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// woker的输出总是要要的
	out := make(chan ParseResult) // 输出
	e.Scheduler.Run()             // 先让Scheduler先Run起来

	// 创建worker执行任务
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler, out)
	}

	// 把seeds扔给Scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	// 拿到处理后的output
	itemCount := 0
	for {
		result := <-out

		// 打印Items结果
		for _, item := range result.Items {
			log.Printf("Got item#%d: %s\n", itemCount, item)
			itemCount++
		}

		// 把新的Requesrts送给Scheduler加进去
		for _, request := range result.Requesrts {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建worker
func createWorker(s Scheduler, out chan ParseResult) {
	// 现在每个worker都有一个产
	// 所以现在chan就是自己的
	in := make(chan Request)
	go func() {
		for {
			// 告诉scheduler我准备好了
			// 把chan作为参数穿进去告诉
			s.WorkerReady(in)
			// 然后呢我们收到事情就做
			request := <-in
			// fetcher网页body(Url+ParseFunc)
			result, err := worker(request)
			if err != nil {
				continue
			}
			// 做完呢就送到out
			out <- result
		}
	}()
}
