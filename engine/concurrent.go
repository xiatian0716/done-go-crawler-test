package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	//任何东西都可以用
	ItemChan chan interface{}
}

type Scheduler interface {
	Submit(Request)
	// 我们要问Scheduler我有一个worker请问给我哪一个chan
	WorkerChan() chan Request
	// WorkerReady(chan Request)
	ReadyNotifier
	// 启动一个总控的goroutine
	Run()
}

// Scheduler里面总共有4个方法，这么大的一个东西送过来有
// 点吃力，我觉得有点重，因此我们把WorkerReady()拿出去
type ReadyNotifier interface {
	// workerChan：从外界告诉我们有一个worker
	// 它可以负责去接收它可以负责去接收request
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// woker的输出总是要要的
	out := make(chan ParseResult) // 输出
	e.Scheduler.Run()             // 先让Scheduler先Run起来

	// 创建worker执行任务
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler, e.Scheduler.WorkerChan(), out)
	}

	// 把seeds扔给Scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	// 拿到处理后的output
	itemCount := 0
	for {
		result := <-out

		// 存储的入口可以叫save
		// 但是我们在这个save里面真的去做事情吗？
		// 真的花时间去save这个item可不可以，这个当然是不可以
		// 我们在这里收到result后，我们作为engine来说它要尽快脱手
		// 它手里不要做很多事情，做事情是worker做的，我们拿到result-item后要尽快脱手
		// save(item)
		// go save(item)
		// go func() { itemChan <- item }()
		// 我们为每个item开一个gorutine之后呢，它们很快就会被消耗掉
		// 消耗的速度比生成的速度快(不需要网络连接-内存)，所以开出来的goruting也不会太多
		// go func() { itemChan <- item }()

		// 打印Items结果
		for _, item := range result.Items {
			// log.Printf("Got item#%d: %s\n", itemCount, item)
			itemCount++
			go func() { e.ItemChan <- item }()
		}

		// 把新的Requesrts送给Scheduler加进去
		for _, request := range result.Requesrts {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建worker
// Scheduler里面总共有4个方法，这么大的一个东西送过来有点吃力，我觉得有点重，因此我们把WorkerReady()拿出去
func createWorker(ready ReadyNotifier, in chan Request, out chan ParseResult) {
	go func() {
		for {
			// 告诉scheduler我准备好了
			// 把chan作为参数穿进去告诉
			ready.WorkerReady(in)
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
