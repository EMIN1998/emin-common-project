package main

import (
	"encoding/json"
	"fmt"
)

// func saveToFile() {
// 	pwd, _ := os.Getwd()
// 	targetPath := pwd + `/file/result.xlsx`
//
// 	// 创建文件，并插入title
// 	targetTitle := []string{"activityId", "language", "reason"}
// 	xlsxFile := getXlsxFile(targetPath, targetTitle)
// 	xlsxSheet := xlsxFile.Sheets[0]
//
// 	// insertRow(xlsxSheet, &targetTitle)
//
// 	// xlsPath := pwd + `\test.xls`
// 	// xlsFile, closer, err := xls.OpenWithCloser(xlsPath, "utf-8")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
//
// 	// 反序列化导入数据
// 	res := make([]*saveData, 0)
// 	data := make([]string, 0)
//
// 	data = append(data, data1, data2, data3, data4, data5, data6, data7, data8, data9, data10, data11)
// 	num := 0
// 	for _, str := range data {
// 		dt := make([]*saveData, 0)
// 		err := json.Unmarshal([]byte(str), &dt)
// 		if err != nil {
// 			fmt.Printf("unmarshal err :%s", err.Error())
// 		}
// 		num += len(dt)
// 		fmt.Printf("\n========> data len %d", len(dt))
// 		res = append(res, dt...)
// 	}
//
// 	fmt.Printf("\n------------------> all len :%d", len(res))
//
// 	// 获取xls文件的第一个sheet
// 	// sheet := xlsFile.GetSheet(0)
// 	// 从第二行开始，遍历xls文件，然后按行调用insertRowFromXls函数
// 	// for j := 1; j < int(sheet.MaxRow)+1; j++ {
// 	// 	xlsRow := sheet.Row(j)
// 	// 	rowColCount := xlsRow.LastCol()
// 	// 	insertData(xlsxSheet, xlsRow, rowColCount)
// 	// }
//
// 	// 插入数据
// 	for _, v := range res {
// 		insertData(xlsxSheet, v.ActivityId, "zh_TW", v.Reason)
// 	}
// 	// closer.Close()
//
// 	// 保存文件
// 	err := xlsxFile.Save(targetPath)
// 	if err != nil {
// 		fmt.Printf("xlsxFile.Save err :%s", err.Error())
// 	}
// }

// 将xls.Row指针对应的数据插入到xlsx.sheet中，待优化
// func insertData(sheet *xlsx.Sheet, activityID int64, language string, reason string) {
// 	row := sheet.AddRow()
// 	// for i := 0; i < rowColCount; i++ {
// 	cell := row.AddCell()
// 	idStr := strconv.FormatInt(activityID, 10)
// 	cell.Value = idStr
//
// 	cellLang := row.AddCell()
// 	cellLang.Value = language
//
// 	cellReason := row.AddCell()
// 	cellReason.Value = reason
// 	// }
//
// }
//
// // 将一个切片指针对应的数据插入到xlsx.sheet中
// func insertRow(sheet *xlsx.Sheet, rowDataPtr *[]string) {
// 	row := sheet.AddRow()
// 	rowData := *rowDataPtr
// 	for _, v := range rowData {
// 		cell := row.AddCell()
// 		cell.Value = v
// 	}
//
// }
//
// // 获取xlsx.File对象的指针，如果文件路径不存在则新建一个文件，并返回其指针
// func getXlsxFile(filePath string, xlsxTitle []string) *xlsx.File {
// 	var file *xlsx.File
// 	if _, err := os.Stat(filePath); err == nil {
// 		file, err = xlsx.OpenFile(filePath)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 	} else {
// 		file = xlsx.NewFile()
// 		sheet, err := file.AddSheet("sheet1")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		insertRow(sheet, &xlsxTitle)
// 	}
//
// 	return file
//
// }
//
// // var statusList = []string{"None", "Done", "Pending"}
//
// func readXslx() []*ReadData {
// 	xlFile, err := xlsx.OpenFile("file/main/resource/china2.xlsx")
// 	if err != nil {
// 		fmt.Print(fmt.Errorf("open file err:%v", err))
// 		return nil
// 	}
//
// 	var (
// 		length     = len(xlFile.Sheets[0].Rows)
// 		sheet      = xlFile.Sheets[0]
// 		valuesList = make([][]string, 0, length)
// 	)
//
// 	// res := make([]*ReadData, 0)
// 	for i, row := range sheet.Rows {
// 		if i == 0 {
// 			continue
// 		}
//
// 		var values []string
// 		// 遍历每一个单元
// 		nilRow := true
// 		for _, cell := range row.Cells {
// 			text := cell.Value
// 			values = append(values, text)
// 			if len(text) > 0 {
// 				nilRow = false
// 			}
// 		}
// 		if nilRow {
// 			continue
// 		}
// 		valuesList = append(valuesList, values)
// 	}
//
// 	if len(valuesList) == 0 {
// 		print(fmt.Errorf("行数有误"))
// 		return nil
// 	}
// 	fmt.Printf("excel 行数：%d", len(valuesList))
//
// 	item := make([]*ReadData, 0)
// 	for i, r := range valuesList {
// 		if len(r) < 4 {
// 			fmt.Printf("数据列数不对，excel行：%d,id ：%v", i+1, r)
// 			continue
// 		}
//
// 		var (
// 			// actId        = 0
// 			// langauge       = ""
// 			// twTargetStatus = int64(0)
// 			// cnTargetStatus = int64(0)
// 			hkTargetStatus = int64(0)
// 		)
//
// 		id, err := strconv.ParseInt(r[1], 10, 64)
// 		if err != nil {
// 			continue
// 		}
//
// 		if !isValidTranslateStatusStr(r[2]) {
// 			continue
// 		}
//
// 		if !isValidTranslateStatusStr(r[3]) {
// 			continue
// 		}
//
// 		if !isValidTranslateStatusStr(r[4]) {
// 			continue
// 		}
//
// 		// twTargetStatus = getTranslateStatus(r[2])
// 		// cnTargetStatus = getTranslateStatus(r[3])
// 		hkTargetStatus = getTranslateStatus(r[4])
//
// 		unit := &ReadData{
// 			"zh_HK",
// 			id,
// 			hkTargetStatus,
// 		}
//
// 		// unitCn := &ReadData{
// 		// 	"zh_CN",
// 		// 	id,
// 		// 	cnTargetStatus,
// 		// }
// 		//
// 		// unitHk := &ReadData{
// 		// 	"zh_HK",
// 		// 	id,
// 		// 	hkTargetStatus,
// 		// }
//
// 		item = append(item, unit)
// 	}
//
// 	bt, err := json.Marshal(item)
// 	if err != nil {
// 		return nil
// 	}
// 	fmt.Printf(string(bt))
// 	return item
// }

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
