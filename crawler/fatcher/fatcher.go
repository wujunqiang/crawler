package fatcher

import (
	"bufio"
	"crawler_company/crawler/httpClient"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

/**
 *通用网络请求器 返回请求的HTML页面
 *
 */

var (
	rateLimiter = time.Tick(
		time.Second * 2)
	verboseLogging = true
)

func SetVerboseLogging() {
	verboseLogging = true
}

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	if verboseLogging {
		log.Printf("Fetching url %s", url)
	}
	//获取URL页面
	//resp, err := httpClient.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//获取cookie
	//cookies := engine.Cookies{}

	resp, err := httpClient.GetDefClientDo(url)
	if err != nil {
		fmt.Errorf("wrong status code: %d",
			resp.StatusCode)
	}

	defer resp.Body.Close()
	//判断访问是否成功。
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}
	//log.Printf("response Header : %s", resp.Header)
	//log.Printf("response Header : %d", resp.)
	//cookies := resp.Header.Get("Cookie")
	//
	//log.Printf("response code : %s", resp.Body)

	//for _, n := range cookies {
	//	log.Println("cookie : " + n.Name + "--" + n.Value)
	//}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

/**
 *转码UTF-8
 */
func determineEncoding(
	r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
