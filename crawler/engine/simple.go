package engine

import (
	"log"
)

/**
 * 调度器
 */

const cookiefile = "/Users/zhangzhengfang/go/src/crawler_company/crawler/config/cookieData.txt"

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {

	var requests []Request
	//便历 种子内的数据
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//判断是不是最后一个
	for len(requests) > 0 {
		r := requests[0]
		//移出一个request
		requests = requests[1:]
		log.Printf("Got Url: %s \n", r.Url)
		parseResult, e := Worker(r)
		if e != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item: %s \n", item)

		}

	}
}
