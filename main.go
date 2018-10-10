package main

import (
	"go-crawler-zhenai/engine"
	"go-crawler-zhenai/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
