package main

import (
	"encoding/json"
	"fmt"
	_ "strconv"
)

func main() {

	res := splitArray(dataTW, 500)
	dataList := make([]*RequestParam, 0)
	//eurLang := []string{"ru_RU", "fr_FR", "it_IT", "es_ES", "de_DE"}
	for _, thId := range res {
		tmp := &RequestParam{}
		tmp.Language = "zh_TW"
		tmp.ActivityIds = JoinInt642(",", thId...)
		dataList = append(dataList, tmp)
	}

	dt, _ := json.Marshal(dataList)
	fmt.Print(string(dt))
}

