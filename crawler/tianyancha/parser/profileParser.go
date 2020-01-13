package parser

import (
	"crawler_company/crawler/engine"
	"crawler_company/crawler/modle"
	"regexp"
)

var NameRE = regexp.MustCompile(`<h1 class="name">([^<]+)</h1>`)
func ParserProfile(content []byte,name string) engine.Request{
	profile :=modle.Profile{}
	profile.Name = name

	profile.Telephone = extractString(content,NameRE)

	return
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}