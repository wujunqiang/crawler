package parser

import (
	"crawler_company/crawler/config"
	"crawler_company/crawler/engine"
	"fmt"
	"regexp"
	"time"
)

//电话
var PhoneRE = regexp.MustCompile(`<span class="label">电话：</span><span>([^<]+)</span>`)

//邮箱
var MailRE = regexp.MustCompile(`>([a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+)<`)

//网址
var WebRE = regexp.MustCompile(`<span class="label">网址：</span><span>([^<]+)</span><`)

//地址
var AddressRE = regexp.MustCompile(`<td width="144px">注册地址</td><td colspan="4">([^<]+)<`)

//法人
var LagerEntiyRE = regexp.MustCompile(`onclick="common.stopPropagation\(event\)">([^<]+)</a>`)

//注册资本
var ZCZB_RE = regexp.MustCompile(`<td width="144px">注册资本</td><td width="308px"><div title="[^"]+">([^<]+)</div>`)

//经营状态
var JYZT_RE = regexp.MustCompile(`<td width="150px">经营状态</td><td width="">([^<]+)<`)

//成立时间
var CLSJ_RE = regexp.MustCompile(`<td width="144px">成立日期</td><td width="308px"><div title=" ">([^<]+)</div>`)

//社会信用代码
var SHXYDM_RE = regexp.MustCompile(`<td width="144px">统一社会信用代码</td><td width="308px">([^<]+)</td>`)

//工商注册号
var GSZCH_RE = regexp.MustCompile(`<td width="150px">工商注册号</td><td>([^<]+)</td>`)

//组织机构代码
var ZZJGDM_RE = regexp.MustCompile(`<td width="150px">组织机构代码</td><td colspan="2">([^<]+)</td>`)

//纳税人实别号
var NSRSBH_RE = regexp.MustCompile(`<td width="144px">纳税人识别号</td><td width="308px">([^<]+)</td>`)

//公司类型
var CompanyType_RE = regexp.MustCompile(`<td width="144px">公司类型</td><td width="308px">([^<]+)</td>`)

//行业
var CompanyHY_RE = regexp.MustCompile(`<td width="150px">行业</td><td colspan="2">([^<]+)</td>`)

//核准日期
var HZRQ_RE = regexp.MustCompile(`<td width="144px">核准日期</td><td
      width="308px">([^<]+)</td>`)

//登记机关
var DJJG_RE = regexp.MustCompile(`<td width="150px">登记机关</td><td colspan="2">([^<]+)</td>`)

//曾用名
var OtherNameRE = regexp.MustCompile(`<td width="144px">曾用名</td><td width="308px"><span>([^<]+)</span></td>`)

//英文名称
var EnglistNameRE = regexp.MustCompile(`<td width="150px">英文名称</td><td colspan="2">([^<]+)</td>`)

//经营范围
var BussinesRE = regexp.MustCompile(`<td width="144px">经营范围</td><td colspan="4"><span class="">([^<]+)</span></td>`)

//ID
var idUrlRe = regexp.MustCompile(
	`https://www.tianyancha.com/company/([\d]+)`)

func ParserProfile(content []byte, url string, name string) engine.ParseResult {

	testurl := extractString(content, testUrlRe)

	result := engine.ParseResult{}
	if testurl == "天眼查校验" {
		fmt.Println("Profile Parser !天眼查校验")
		engine.SetDuplicat(url)
		time.Sleep(time.Duration(30) * time.Second)
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    url,
				Parser: NewProfileParser(name),
			})
		return result

	} else {

		profile := make([]string, 20)
		profile[0] = name

		profile[1] = extractString(content, PhoneRE)

		profile[2] = extractString(content, MailRE)

		profile[3] = extractString(content, WebRE)

		profile[4] = extractString(content, LagerEntiyRE)

		profile[5] = extractString(content, AddressRE)

		profile[6] = extractString(content, ZCZB_RE)

		profile[7] = extractString(content, JYZT_RE)

		profile[8] = extractString(content, CLSJ_RE)

		profile[9] = extractString(content, SHXYDM_RE)

		profile[10] = extractString(content, GSZCH_RE)

		profile[11] = extractString(content, ZZJGDM_RE)

		profile[12] = extractString(content, NSRSBH_RE)

		profile[13] = extractString(content, CompanyType_RE)

		profile[14] = extractString(content, CompanyHY_RE)

		profile[15] = extractString(content, HZRQ_RE)

		profile[16] = extractString(content, DJJG_RE)

		profile[17] = extractString(content, EnglistNameRE)

		profile[18] = extractString(content, OtherNameRE)

		profile[19] = extractString(content, BussinesRE)

		result = engine.ParseResult{
			Items: []engine.Item{
				{
					Url:  url,
					Type: "tianyancha",
					Id: extractString(
						[]byte(url), idUrlRe),
					Payload: profile,
				},
			},
		}
	}
	return result
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

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(
	contents []byte,
	url string) engine.ParseResult {
	return ParserProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (
	name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(
	name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
