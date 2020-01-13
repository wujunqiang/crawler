package parser

import "testing"

func TestCompany(t *testing.T) {
	ParseCompany([]byte(`<a class="name " tyc-event-click="" tyc-event-ch="CompanySearch.Company" href="https://www.tianyancha.com/company/3114495038" target="_blank">中国海洋石油集团有限公司</a>`))
}