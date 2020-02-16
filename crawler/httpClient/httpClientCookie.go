package httpClient

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Unknwon/goconfig"

	"github.com/onyas/go-browsercookie"
)

const cookiefile = "/Users/zhangzhengfang/go/src/crawler_company/crawler/config/cookieData.txt"

//var client *http.Client
//var req http.Request

func GetDefClientDo(url string) (*http.Response, error) {

	c, _ := goconfig.LoadConfigFile("conf.ini")
	u, _ := c.GetValue("page", "url")

	//获取chrome浏览器的Cookie
	cookieJar, _ := browsercookie.Chrome(u)

	//fmt.Println("CookieJar=", cookieJar)
	client := &http.Client{Jar: cookieJar}

	//准备请求地址，和方法
	req, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return nil, e
	}

	//for _, m := range getCookies() {
	//
	//	req.AddCookie(&m)
	//}

	//模仿Chrome设置请求头
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "text/html, charset=utf-8")
	req.Header.Set("Transfer-Encoding", "chunked")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	//req.Header.Add("cookie", getCookiestr())

	//fmt.Printf("data is : %s ",req.Header)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("cookies response : %v \n", response.Header.Get("cookie"))

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
		//fmt.Println(str[0] + " = " + str[1])
		cookie := http.Cookie{Name: str[0], Value: str[1]}
		cookies = append(cookies, cookie)
	}

	return cookies
}
