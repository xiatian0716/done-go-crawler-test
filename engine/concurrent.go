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
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// 所有的woker共用一个输入
	in := make(chan Request)                  // 输入
	e.Scheduler.ConfigureMasterWorkerChan(in) // 把in放到Scheduler里面
	out := make(chan ParseResult)             // 输出

	// 创建worker执行任务
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
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
			log.Printf("Got item#%d: %s\n", itemCount,item)
			itemCount++
		}

		// 把Requesrts送给Scheduler
		for _, request := range result.Requesrts {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建worker
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			// fetcher网页body(Url+ParseFunc)
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
