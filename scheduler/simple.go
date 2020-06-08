package scheduler

import "go-crawler-test/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 他问我要WorkerChan我就return这个workerChan
// 所以这样所有的worker就共用这一个workerChan了
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// WorkerReady不做事情，但是我们也实现一下，这样就实现了这个接口
func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
}

// 在Run里面我们就把workerChan做出来
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//把request发送给worker chan
func (s *SimpleScheduler) Submit(request engine.Request) {
	// s.workerChan <- request
	go func() { s.workerChan <- request }()
}
