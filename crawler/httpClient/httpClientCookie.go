package httpClient

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const cookiefile = "/Users/zhangzhengfang/go/src/crawler_company/crawler/config/cookieData.txt"

var client *http.Client
var req http.Request

func GetDefClientDo(url string) (*http.Response, error) {

	client = &http.Client{}

	//准备请求地址，和方法
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return nil, e
	}

	for _, m := range getCookies() {
		req.AddCookie(&m)
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Encoding", "gzip")
	req.Header.Add("Content-Type", "text/html, charset=utf-8")
	req.Header.Add("Transfer-Encoding", "chunked")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("cookie",getCookiestr())

	//fmt.Printf("data is : %s ",req.Header)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getCookiestr() string {

	//读取本地cookie文件。合并成一组cookie数据，传递到下一层
	bytes, err := ioutil.ReadFile(cookiefile)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func getCookies() []http.Cookie {

	//读取本地cookie文件。合并成一组cookie数据，传递到下一层
	bytes, err := ioutil.ReadFile(cookiefile)
	if err != nil {
		panic(err)
	}
	cookies := []http.Cookie{}
	splitstr := strings.Split(string(bytes), "; ")
	for _, m := range splitstr {
		str := strings.Split(m, "=")
		cookie := http.Cookie{Name: str[0], Value: str[1]}
		cookies = append(cookies, cookie)
	}

	return cookies
}
