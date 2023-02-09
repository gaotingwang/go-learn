package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	ce.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		// 相当于 in <- r
		ce.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: %v\n", item)
		}
		for _, request := range result.Requests {
			// 相当于 in <- r
			// out产生的结果等待写入in ，in中数据等待out处理完，产生了阻塞，在SimpleScheduler.Submit为每个request开启一个goroutine
			ce.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
