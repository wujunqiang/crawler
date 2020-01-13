package parser

import (
	"crawler_company/crawler/engine"
	"fmt"
	"regexp"
)

const companyRe  =`href="(https://www.tianyancha.com/company/[0-9]*)+"*[^>]*>([^<]+)</a>`
func ParseCompany(content []byte) engine.ParseResult {
	re := regexp.MustCompile(companyRe)
	submatch := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _,m := range submatch{
		fmt.Printf("URL: %s , COMPANY: %s \n",m[1],m[2])
		result.Items = append(result.Items,m[2])
		result.Requests  = append( result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	//fmt.Printf("%d",len(submatch))
	return result
}