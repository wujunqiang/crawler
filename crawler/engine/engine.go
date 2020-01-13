package engine

import (
	"crawler_company/crawler/fatcher"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
 * 调度器
 */

const cookiefile  = "/Users/zhangzhengfang/go/src/crawler_company/crawler/config/cookieData.txt"

func Run(seeds ...Request)  {

	//读取本地cookie文件。合并成一组cookie数据，传递到下一层
	bytes, err := ioutil.ReadFile(cookiefile)
	if err != nil{
		panic(err)
	}
	cookies := []http.Cookie{}
	splitstr := strings.Split(string(bytes), "; ")
	for _,m := range splitstr{
		str := strings.Split(m, "=")
		log.Printf("this is cookie name  and value %s " ,str)
		cookie := http.Cookie{Name:str[0],Value: str[1]}
		cookies = append(cookies,cookie)
	}

	var requests []Request
	//便历 种子内的数据
	for _, r := range seeds{
		requests = append(requests, r)
	}
	//判断是不是最后一个
	for len(requests)  > 0{
		r := requests[0]
	//移出一个request
		requests = requests[1:]
		log.Printf("Got Url: %s \n" ,r.Url)
		body, e := fatcher.Fetch(r.Url,cookies)
		if e != nil {
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items{
			log.Printf("Got item: %s \n" ,item)


		}

	}
}