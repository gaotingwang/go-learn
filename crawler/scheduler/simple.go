package scheduler

import (
	"crawler/engine"
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

func (s *SimpleScheduler) ConfigureMasterWorkerChan(rc chan engine.Request) {
	s.workerChan = rc
}
