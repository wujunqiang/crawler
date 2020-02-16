package parser

import (
	"crawler_company/crawler/config"
	"crawler_company/crawler/engine"

	"regexp"
)

const cityListRe = `<a href="(https://www.tianyancha.com/search+[?][^"]*)+"[^>]*>([^<]+)</a>`

var areaListRe = regexp.MustCompile(`<a class="item" href="(https://www.tianyancha.com/search?base=[a-z]+&amp;areaCode=[^"]+)" href-new-event="" event-name=[^<]+</a>`)

func ParseCityList(content []byte, _ string) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range submatch {
		//fmt.Printf("URL: %s , CITY: %s \n",m[1],m[2])

		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(
				ParseCompany, config.ParseCompany),
		})
	}
	//fmt.Printf("%d",len(submatch))
	return result
}
