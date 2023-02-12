package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(request Request)
	WorkerChan() chan Request
	Run()
	ReadyNotify
}

type ReadyNotify interface {
	WorkerReady(chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	ce.Scheduler.Run()

	for i := 0; i < ce.WorkerCount; i++ {
		// 由scheduler决定worker输入channel
		// simple 的输入channel为所有worker共用，queued 的输入channel是为每个worker创建单独的chan Request
		ce.createWorker(ce.Scheduler.WorkerChan(), out, ce.Scheduler)
	}

	for _, r := range seeds {
		// 相当于 in <- r
		ce.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				ce.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			// 相当于 in <- r
			// 在simpleScheduler情况下，out产生的结果等待写入in ，in中数据等待out处理完，产生了阻塞，需要在SimpleScheduler.Submit为每个request开启一个goroutine
			ce.Scheduler.Submit(request)
		}
	}
}

// worker 从in获取数据，处理完结果写入out中
func (ce *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, notify ReadyNotify) {
	go func() {
		for {
			notify.WorkerReady(in)
			request := <-in
			result, err := ce.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
