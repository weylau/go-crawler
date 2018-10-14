package engine

import "fmt"

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurentEngine struct {
	Scheduler Scheduler
	WorkerCnt int
}

func (e *ConcurentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCnt; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: %v\n", item)
		}

		for _, request := range result.Results {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
