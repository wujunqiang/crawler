package persist

import (
	"crawler_company/crawler/engine"
	"github.com/tealeg/xlsx"
	"log"
)

const xlxsname = "/Users/zhangzhengfang/go/src/crawler/outEngineData.xlsx"

func ItemSaver(index string) (chan engine.Item, error) {
	_, err := xlsx.OpenFile(xlxsname)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

		}
	}()

	return out, nil

}
