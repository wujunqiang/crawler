package engine

import (
	"fmt"
	"log"

	. "github.com/Unknwon/goconfig"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
	Cgf              *ConfigFile
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(e.Scheduler.WorkerChan(),
			out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func (e *ConcurrentEngine) CreateWorker(
	in chan Request, out chan ParseResult,
	ready ReadNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				fmt.Printf("The crteat work is err : %v", err)
				continue
			}
			out <- result
		}
	}()
}

var VisitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	log.Printf("duplicate vistiedUrl lens : %d", len(VisitedUrls))
	if VisitedUrls[url] {
		return true
	}

	VisitedUrls[url] = true
	return false

}

//检查打码时，重置拦截网址为false
func SetDuplicat(url string) {
	VisitedUrls[url] = false

}
