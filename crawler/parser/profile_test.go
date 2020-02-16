package parser

import (
	"fmt"
	"testing"
)

func TestParserProfile(t *testing.T) {

	ints := make([]int, 4)
	fmt.Println(len(ints))
	ints = ints[1:]
	fmt.Println(len(ints))

	//contents, err := ioutil.ReadFile(
	//	"/Users/zhangzhengfang/go/src/crawler/crawler/config/profileData.html")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//result := ParserProfile(contents,
	//	"乌海岱山林牧业有限公司",
	//	"https://www.tianyancha.com/company/294126957",
	//)
	//
	//if len(result.Items) != 1 {
	//	t.Errorf("Items should contain 1 "+
	//		"element; but was %v", result.Items)
	//}
	//
	//actual := result.Items[0]
	//
	//expected := engine.Item{
	//	Url:  "https://www.tianyancha.com/company/294126957",
	//	Type: "tianyancha",
	//	Id:   "294126957",
	//	Payload: modle.Profile{
	//		Name:        "乌海岱山林牧业有限公司",
	//		Telephone:   "13947321809",
	//		WEB:         "暂无信息",
	//		ADDRESS:     "内蒙古自治区乌海市海区摩尔沟东街B-5#",
	//		EMail:       "wuhaidaishan@163.com",
	//		LagelEntiy:  "杜仲义",
	//		ZCZB:        "12000万",
	//		JYZT:        "存续",
	//		CLSJ:        "2002-06-13",
	//		SHXYDM:      "91150300740112796J",
	//		GSZCH:       "150300000000061",
	//		ZZJGDM:      "740112796",
	//		NSRSBH:      "91150300740112796J",
	//		CompanyType: "有限责任公司(自然人投资或控股)(1130)",
	//		CompanyHY:   "林业",
	//		HZRQ:        "2018-01-29",
	//		DJJG:        "乌海市工商行政管理局",
	//		OtherName:   "-",
	//		EnglistName: "-",
	//	},
	//}
	//
	//if actual != expected {
	//	t.Errorf("expected %v\n;                but was %v",
	//		expected, actual)
	//}
}
