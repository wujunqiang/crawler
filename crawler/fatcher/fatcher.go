package fatcher

import (
	"bufio"
	"crawler_company/crawler/httpClient"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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
//qps
var (
	rateLimiter = time.Tick(
		time.Second / 3)
	verboseLogging = true
)

//随机sleep 实测无效
func radomTime() time.Duration {
	i := rand.Intn(3)
	duration := time.Duration(i) * time.Second
	return duration
}

func SetVerboseLogging() {
	verboseLogging = true
}

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	if verboseLogging {
		log.Printf("Fetching url %s", url)
	}

	resp, err := httpClient.GetDefClientDo(url)
	if err != nil {
		fmt.Errorf("wrong status code: %d",
			err)

		//fmt.Println("URL ：" + url)

		return []byte{}, nil
	}

	defer resp.Body.Close()
	//判断访问是否成功。
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}
	//log.Printf("response Header : %s", resp.Body)
	//log.Printf("response Header : %d", resp.Cookies())
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
	//log.Printf("response Header : %s", utf8Reader)
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
