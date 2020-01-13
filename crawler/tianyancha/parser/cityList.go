package parser

import (
	"crawler_company/crawler/engine"

	"regexp"
)

const cityListRe  = `<a href="(https://www.tianyancha.com/search+[?][^"]*)+"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult   {


	re := regexp.MustCompile(cityListRe)
	submatch := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _,m := range submatch{
		//fmt.Printf("URL: %s , CITY: %s \n",m[1],m[2])
		result.Items = append(result.Items,m[2])
		result.Requests  = append( result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCompany,
		})
	}
	//fmt.Printf("%d",len(submatch))
	return result
}
