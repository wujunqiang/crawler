package main

import (
	"crawler_company/crawler/engine"
	"crawler_company/crawler/tianyancha/parser"
	"go.uber.org/zap"
)

/*
 *天眼查有登陆机制。为了访问下一层。每次启动前，在WEB上登陆后，拿到COOKIE，存放在本地cookieData.txt文件中
 */

func main() {

	logger, _ := zap.NewProduction()
	logger.Info("this is running")
	engine.Run(engine.Request{
		Url:        "https://www.tianyancha.com/",
		ParserFunc: parser.ParseCityList,
	})

}
