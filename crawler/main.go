package main

import (
	"crawler_company/crawler/config"
	"crawler_company/crawler/engine"
	"crawler_company/crawler/persist"
	"crawler_company/crawler/scheduler"
	"crawler_company/crawler/parser"
)

/*
 *天眼查有登陆机制。为了访问下一层。每次启动前，在WEB上登陆后，拿到COOKIE，存放在本地cookieData.txt文件中
 */

func main() {

	//logger, _ := zap.NewProduction()
	//logger.Info("this is running")
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "https://www.tianyancha.com/",
	//	Parser: engine.NewFuncParser(
	//		parser.ParseCityList,config.ParseCityList),
	//})
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:			&scheduler.QueuedScheduler{},
		WorkerCount:		5,
		ItemChan:			itemChan,
		RequestProcessor:	engine.Worker,
	}

	e.Run(engine.Request{
		Url:	 "https://www.tianyancha.com/",
		Parser:	engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})

}
