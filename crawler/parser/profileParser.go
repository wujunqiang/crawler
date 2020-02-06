package parser

import (
	"crawler_company/crawler/config"
	"crawler_company/crawler/engine"
	"crawler_company/crawler/modle"
	"regexp"
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
var ZCZB_RE = regexp.MustCompile(`<td width="144px">注册资本</td><td width="308px"><div title="[0-9]+万">([^<]+)</div>`)

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
var OtherNameRE = regexp.MustCompile(`<td width="144px">曾用名</td><td width="308px">([^<]+)</td>`)

//英文名称
var EnglistNameRE = regexp.MustCompile(`<td width="150px">英文名称</td><td colspan="2">([^<]+)</td>`)

//ID
var idUrlRe = regexp.MustCompile(
	`https://www.tianyancha.com/company/([\d]+)`)

func ParserProfile(content []byte, name string, url string) engine.ParseResult {
	profile := modle.Profile{}
	profile.Name = name

	profile.Telephone = extractString(content, PhoneRE)

	profile.EMail = extractString(content, MailRE)

	profile.WEB = extractString(content, WebRE)

	profile.LagelEntiy = extractString(content, LagerEntiyRE)

	profile.ADDRESS = extractString(content, AddressRE)

	profile.ZCZB = extractString(content, ZCZB_RE)

	profile.JYZT = extractString(content, JYZT_RE)

	profile.CLSJ = extractString(content, CLSJ_RE)

	profile.SHXYDM = extractString(content, SHXYDM_RE)

	profile.GSZCH = extractString(content, GSZCH_RE)

	profile.ZZJGDM = extractString(content, ZZJGDM_RE)

	profile.NSRSBH = extractString(content, NSRSBH_RE)

	profile.CompanyType = extractString(content, CompanyType_RE)

	profile.CompanyHY = extractString(content, CompanyHY_RE)

	profile.HZRQ = extractString(content, HZRQ_RE)

	profile.DJJG = extractString(content, DJJG_RE)

	profile.EnglistName = extractString(content, EnglistNameRE)

	profile.OtherName = extractString(content, OtherNameRE)

	result := engine.ParseResult{
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
