package parser

import (
	"crawler_company/crawler/config"
	"crawler_company/crawler/engine"
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"github.com/Unknwon/goconfig"
)

var companyRe = regexp.MustCompile(`href="(https://www.tianyancha.com/company/[0-9]*)+"*[^>]*>([^<]+)</a>`)

var companyPageRe = regexp.MustCompile(`<a class="num -next"[^\s]*([https]*[^"]*[^>]*)>`)

var urlRe = regexp.MustCompile(`([a-zA-z]+://[^\s^"]*)`)

var testUrlRe = regexp.MustCompile("<title>([^<]+)</title>")

const CONFIGFILE = "conf.ini"

func ParseCompany(content []byte, _ string) engine.ParseResult {

	testurl := extractString(content, testUrlRe)

	result := engine.ParseResult{}

	Cfg, err := goconfig.LoadConfigFile(CONFIGFILE)
	if err != nil {
		fmt.Errorf("this config file not find %v", err)
	}

	if testurl == "天眼查校验" {
		locatUrl, err := Cfg.GetValue("page", "url")
		if err != nil {
			panic("错误")
		}
		fmt.Println("Company Parser !天眼查校验")
		//log.Printf("wait 30 second time begin! : %s", locatUrl)
		//执行chrome浏览器操作打码
		exec.Command(`open`, locatUrl).Start()
		engine.SetDuplicat(locatUrl)
		time.Sleep(time.Duration(20) * time.Second)
		//放入config保存地址到队列
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(locatUrl),
				Parser: engine.NewFuncParser(
					ParseCompany, config.ParseCompany),
			})

		//log.Println("wait 30 second time end!")
		return result //此处不返回，result无法送回队列。原因不知道。

	} else {
		submatch := companyRe.FindAllSubmatch(content, -1)

		companyUrl := extractString(content, companyPageRe)

		cUrl := extractString([]byte(companyUrl), urlRe)

		for _, m := range submatch {
			fmt.Printf("URL: %s , COMPANY: %s \n", m[1], m[2])

			result.Requests = append(
				result.Requests, engine.Request{
					Url:    string(m[1]),
					Parser: NewProfileParser(string(m[2])),
				})
		}

		//存储翻页的URL到config文件里
		Cfg.SetValue("page", "url", cUrl)
		goconfig.SaveConfigFile(Cfg, CONFIGFILE)
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(cUrl),
				Parser: engine.NewFuncParser(
					ParseCompany, config.ParseCompany),
			})
	}

	return result
}
