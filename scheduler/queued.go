package scheduler

import "go-crawler-test/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// 给worker返回一个chan
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	// 我们希望每个workerChan有一个自己的chan
	return make(chan engine.Request)
}

// requestChan：有人Submit一个request我们就加进去
func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

// workerChan：从外界告诉我们有一个worker它可以负责去接收它可以负责去接收request
func (s *QueuedScheduler) WorkerReady(workerChan chan engine.Request) {
	s.workerChan <- workerChan
}

// 启动一个总控的goroutine
func (s *QueuedScheduler) Run() {
	// 因为要生成它们，我们改变了s的内容，所以都要改成
	// 指针(*)接受者，指针(*)接受者才能改变里面的内容
	// 生成requestChan-workerChan开始做事情
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		// 不断的做事情
		for {
			// 如果我们既有request在排队
			// 又有worker在排队就可以发
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			// select两件独立的事情
			select {
			// request从requestChan里拿
			case request := <-s.requestChan:
				// 收到request就让request排队
				requestQ = append(requestQ, request)
			// worker从workerChan里拿
			case worker := <-s.workerChan:
				// 收到worker就让worker排队
				workerQ = append(workerQ, worker)
			// 什么情况下我们可以把request发给worker呢？
			case activeWorker <- activeRequest:
				// 送完后把它们从队列里面拿掉
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
