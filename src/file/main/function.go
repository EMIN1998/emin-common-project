package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/klook/klook-libs/src/klook.libs/locale"
	"github.com/tealeg/xlsx"
)

func saveToFile(title []string, resource []*SaveData, fileName string) error {
	pwd, _ := os.Getwd()
	targetPath := fmt.Sprintf("%s/file/main/resource/%s.xlsx", pwd, fileName)

	if len(resource) == 0 || resource[0] == nil {
		return nil
	}

	// 创建文件，并插入title
	xlsxFile := getXlsxFile(targetPath, title)
	xlsxSheet := xlsxFile.Sheets[0]

	// insertRow(xlsxSheet, &targetTitle)

	// xlsPath := pwd + `\test.xls`
	// xlsFile, closer, err := xls.OpenWithCloser(xlsPath, "utf-8")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 反序列化导入数据
	// res := make([]*saveData, 0)
	// data := make([]string, 0)

	// data = append(data)
	// num := 0
	// for _, str := range data {
	// 	dt := make([]*saveData, 0)
	// 	err := json.Unmarshal([]byte(str), &dt)
	// 	if err != nil {
	// 		fmt.Printf("unmarshal err :%s", err.Error())
	// 	}
	// 	num += len(dt)
	// 	fmt.Printf("\n========> data len %d", len(dt))
	// 	res = append(res, dt...)
	// }

	// fmt.Printf("\n------------------> all len :%d", len(res))

	// 获取xls文件的第一个sheet
	// sheet := xlsFile.GetSheet(0)
	// 从第二行开始，遍历xls文件，然后按行调用insertRowFromXls函数
	// for j := 1; j < int(sheet.MaxRow)+1; j++ {
	// 	xlsRow := sheet.Row(j)
	// 	rowColCount := xlsRow.LastCol()
	// 	insertData(xlsxSheet, xlsRow, rowColCount)
	// }

	// 插入数据
	for _, v := range resource {
		args := make([]string, 0)
		// t := reflect.TypeOf(v)
		args = append(args, v.Sql)

		insertData(xlsxSheet, args)
	}

	// 保存文件
	err := xlsxFile.Save(targetPath)
	if err != nil {
		fmt.Printf("xlsxFile.Save err :%s", err.Error())
		return err
	}

	return nil
}

// 将xls.Row指针对应的数据插入到xlsx.sheet中，待优化
func insertData(sheet *xlsx.Sheet, args []string) {
	row := sheet.AddRow()

	for _, arg := range args {
		cell := row.AddCell()
		value := fmt.Sprintf("%v", arg)
		cell.Value = value
	}

}

// 将一个切片指针对应的数据插入到xlsx.sheet中
func insertRow(sheet *xlsx.Sheet, rowDataPtr *[]string) {
	row := sheet.AddRow()
	rowData := *rowDataPtr
	for _, v := range rowData {
		cell := row.AddCell()
		cell.Value = v
	}

}

// 获取xlsx.File对象的指针，如果文件路径不存在则新建一个文件，并返回其指针
func getXlsxFile(filePath string, xlsxTitle []string) *xlsx.File {

	var file *xlsx.File
	if _, err := os.Stat(filePath); err == nil {
		file, err = xlsx.OpenFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		file = xlsx.NewFile()
		sheet, err := file.AddSheet("sheet1")
		if err != nil {
			log.Fatal(err)
		}
		insertRow(sheet, &xlsxTitle)
	}

	return file

}

// var statusList = []string{"None", "Done", "Pending"}
var langSlice = []string{"en_US",
	"zh_CN",
	"zh_TW",
	"zh_HK",
	"ja_JP",
	"ko_KR",
	"th_TH",
	"id_ID",
	"vi_VN",
	"fr_FR",
	"de_DE",
	"it_IT",
	"es_ES",
	"ru_RU"}

func readXslx() interface{} {
	idxMap := make(map[string]int)
	for idx, lang := range langSlice {
		idxMap[lang] = idx + 2
	}

	xlFile, err := xlsx.OpenFile("file/main/resource/subcategory_no_price_text.xlsx")
	if err != nil {
		fmt.Print(fmt.Errorf("open file err:%v", err))
		return nil
	}

	var (
		length     = len(xlFile.Sheets[0].Rows)
		sheet      = xlFile.Sheets[0]
		valuesList = make([][]string, 0, length)
	)

	res := make([]*NoPriceResource, 0)
	for i, row := range sheet.Rows {
		if i == 0 {
			continue
		}

		var values []string
		// 遍历每一个单元
		nilRow := true
		for _, cell := range row.Cells {
			text := cell.Value
			values = append(values, text)
			if len(text) > 0 {
				nilRow = false
			}
		}
		if nilRow {
			continue
		}
		valuesList = append(valuesList, values)
	}

	if len(valuesList) == 0 {
		print(fmt.Errorf("行数有误"))
		return nil
	}
	fmt.Printf("excel 行数：%d", len(valuesList))

	// item := make(map[string]string, 0)
	for i, r := range valuesList {
		if len(r) < 4 {
			fmt.Printf("数据列数不对，excel行：%d,id ：%v", i+1, r)
			continue
		}

		var (
			// actId        = 0
			// langauge       = ""
			// twTargetStatus = int64(0)
			// cnTargetStatus = int64(0)
			// hkTargetStatus = int64(0)
			textMap = make(map[string]string)
		)

		langList := locale.GetPublishedLangList()
		for _, lang := range langList {
			idx := idxMap[lang]
			if strings.HasPrefix(lang, "en") && lang != locale.EN_US {
				idx = idxMap[locale.EN_US]
			}

			textMap[lang] = r[idx]
		}

		id, err := strconv.ParseInt(r[0], 10, 64)
		if err != nil {
			continue
		}

		textMap[locale.EN_US] = r[2]
		item := NoPriceResource{
			id,
			textMap,
		}

		res = append(res, &item)

	}

	for _, v := range res {
		bt, err := json.Marshal(v)
		if err != nil {
			return nil
		}

		fmt.Printf("\n %s \n", string(bt))
	}

	bt, err := json.Marshal(res)
	if err != nil {
		return nil
	}

	fileName := "test.dat"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer dstFile.Close()
	s := string(bt)
	_, err = dstFile.WriteString(s + "\n")
	if err != nil {
		return err
	}
	return nil
}

func getTranslateStatus(str string) int64 {
	var targetStatus int64
	switch str {
	case "None":
		targetStatus = 0
	case "Pending":
		targetStatus = 1
	case "Done":
		targetStatus = 2
	}
	return targetStatus
}

func isValidTranslateStatusStr(str string) bool {
	if str != "None" && str != "Pending" && str != "Done" {
		return false
	}

	return true
}

func InitDataCleanValue() ([]*ReadData, error) {

	data := make([]*ReadData, 0)
	dataStr := []string{dataStr1, dataStr2, dataStr3, dataStr4}
	for _, str := range dataStr {
		v := make([]*ReadData, 0)
		err := json.Unmarshal([]byte(str), &v)
		if err != nil {
			return nil, err
		}

		for _, unit := range v {
			if unit.Status == 2 {
				data = append(data, unit)
			}
		}
	}

	return data, nil
}

func parseSqlForMfo() ([]*MfoKeyWord, error) {
	data := make([]*MfoKeyWord, 5000)
	err := json.Unmarshal([]byte(mfoString), &data)
	if err != nil {
		fmt.Errorf("json err :%s", err)
		return data, err
	}

	return data, nil
}
