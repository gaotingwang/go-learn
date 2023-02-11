package scheduler

import (
	"github.com/gaotingwang/go-learn/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 所有worker共用一个输入workerChan
func (s *SimpleScheduler) Submit(request engine.Request) {
	// 防止阻塞，为每个request都创建了goroutine
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
