package persist

import (
	"crawler_company/crawler/engine"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var xlsxname = "EngineDate" + time.Now().Format("2006-01-02 15:04:05") + ".xlsx"

var build strings.Builder

func ItemSaver(index string) (chan engine.Item, error) {
	//每次进入重新创建文件
	f := excelize.NewFile()
	if e := f.SaveAs(xlsxname); e != nil {
		panic(e.Error())
	}
	file, e := excelize.OpenFile(xlsxname)
	if e != nil {
		panic(e)
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 1
		for {
			item := <-out
			//log.Printf("Item Saver: got item %d: %v", itemCount, item)
			itemCount++
			xlsx := make(map[string]string)
			temp := 'A'
			for i := 0; i < len(item.Payload); i++ {
				build.WriteString(string(temp))
				build.WriteString(strconv.Itoa(itemCount))
				a := build.String()
				build.Reset()
				xlsx[a] = item.Payload[i]
				temp++
			}
			//log.Println(xlsx)
			save(file, xlsx)

		}
	}()

	return out, nil

}

func save(file *excelize.File, xlsx map[string]string) {

	for k, v := range xlsx {
		//fmt.Printf("%s - %s - %s", "Sheet1", string(k), v)
		file.SetCellValue("Sheet1", string(k), v)
	}

	if e := file.Save(); e != nil {
		panic(e.Error())
	}
}

//func save1f *excelize.File, i map[string]string, item engine.Item) {
//
//	//for i := 1; i < 19; i++ {
//
//	build.WriteString(string(itemCount))
//
//	build.WriteString(strconv.Itoa(i))
//	s := build.String()
//	build.Reset()
//	log.Printf("string is : %s", s)
//	f.SetCellValue("Sheet1", s, item.Payload.Name)
//
//	//}
//	if e := f.Save(); e != nil {
//		panic(e)
//	}
//
//}
