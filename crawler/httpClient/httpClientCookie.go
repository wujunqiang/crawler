package httpClient

import (
	"net/http"
)

func GetDefClient(method string,url string,cookies []http.Cookie) (*http.Response,error) {

	client := &http.Client{}
	var req *http.Request
	//准备请求地址，和方法
	req, e := http.NewRequest(method, url, nil)
	if e != nil {
		return nil ,e
	}

	for _,m := range cookies{
		req.AddCookie(&m)
	}

	//fmt.Printf("data is : %s ",bytes)
	response, err := client.Do(req)
		if err != nil {
			return nil ,err
		}
	return response,nil
}
