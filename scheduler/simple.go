package scheduler

import "go-crawler-test/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 把in放到Scheduler里面
// * ConfigureMasterWorkerChan会改变workerChan的值
func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}

//把request发送给worker chan
func (s *SimpleScheduler) Submit(request engine.Request) {
	// s.workerChan <- request
	go func() { s.workerChan <- request }()
}
