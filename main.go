package main

import (
	"go-crawler/engine"
	"go-crawler/scheduler"
	"go-crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCnt: 10,
	}

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParserCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/changsha",
		ParserFunc: parser.ParserCity,
	})

}
