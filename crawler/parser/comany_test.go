package parser

import (
	"fmt"
	"testing"

	"github.com/Unknwon/goconfig"
)

func TestCompany(t *testing.T) {
	Cfg, err := goconfig.LoadConfigFile("/Users/zhangzhengfang/go/src/crawler/conf.ini")
	if err != nil {
		fmt.Errorf("this config file not find %v", err)
	}

	locatUrl, err := Cfg.GetValue("page", "url")
	if err != nil {
		panic("错误")
	}
	fmt.Println(locatUrl)

	value := Cfg.SetValue("page", "url", "https://www.tianyancha.com/search/p6?base=bj")
	fmt.Println(value)
	goconfig.SaveConfigFile(Cfg, "/Users/zhangzhengfang/go/src/crawler/conf.ini")
	/*
		contents, err := ioutil.ReadFile(
			"/Users/zhangzhengfang/go/src/crawler/crawler/config/profileData.html")

		if err != nil {
			panic(err)
		}
		company := ParseCompany(contents, config.ParseCompany)

		fmt.Printf("companyUrl : %s", company)
	*/
	//----------------------------------------
	//
	//resp, err := http.Get("https://www.tianyancha.com/search?base=sjz")
	//if err != nil{
	//	panic(err)
	//}
	//
	//bytes, e := ioutil.ReadAll(resp.Body)
	//if e != nil{
	//	panic(e)
	//}
	//
	//fmt.Printf("%s",bytes)
	//ParseCompany(bytes,config.ParseCompany)

	//Println("hello")

	//Println(cap(a2))
}
