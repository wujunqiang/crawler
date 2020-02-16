package main

import (
	"crawler_company/crawler/config"
	"crawler_company/crawler/engine"
	"crawler_company/crawler/parser"
	"crawler_company/crawler/persist"
	"crawler_company/crawler/scheduler"
	"fmt"
	"os"

	"github.com/Unknwon/goconfig"
)

/*
 *天眼查有登陆机制。为了访问下一层。每次启动前，在WEB上登陆后，拿到COOKIE，存放在本地cookieData.txt文件中
 */
var CompCfg goconfig.ConfigFile

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

	CompCfg, err2 := goconfig.LoadConfigFile("conf.ini")
	if err2 != nil {

	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      30,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
		Cgf:              CompCfg,
	}

	Begurl, err := e.Cgf.GetValue("page", "url")
	if err != nil {
		fmt.Errorf("load config file err :%v ", err)
	}

	e.Run(engine.Request{
		Url: Begurl,
		Parser: engine.NewFuncParser(
			parser.ParseCompany,
			config.ParseCompany),
	})

}
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
