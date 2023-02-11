package scheduler

import "github.com/gaotingwang/go-learn/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) WorkerReady(rc chan engine.Request) {
	q.workerChan <- rc
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
			case r := <-q.requestChan:
				// send r to a worker
				requestQueue = append(requestQueue, r)
			case w := <-q.workerChan:
				// send w.nextRequest to w
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
