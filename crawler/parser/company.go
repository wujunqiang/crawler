package parser

import (
	"crawler_company/crawler/engine"
	"fmt"
	"regexp"
)

const companyRe = `href="(https://www.tianyancha.com/company/[0-9]*)+"*[^>]*>([^<]+)</a>`

func ParseCompany(content []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(companyRe)
	submatch := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range submatch {
		fmt.Printf("URL: %s , COMPANY: %s \n", m[1], m[2])

		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}
	//fmt.Printf("profiler num is :%d",len(result.Requests))
	return result
}
