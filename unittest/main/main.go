package main

import "time"

// func readDataFromXlsx() ([]*UpdateLanguageStatusBatchReq, error) {
// 	importItems := make([]*UpdateLanguageStatusBatchReq, 0)
// 	xlFile, err := xlsx.OpenFile("apis/activities/activityimport/resource/chinatown_automation.xlsx")
// 	if err != nil {
// 		return nil, errors.Errorf(err, "open file err")
// 	}
//
// 	var (
// 		length     = len(xlFile.Sheets[0].Rows)
// 		sheet      = xlFile.Sheets[0]
// 		valuesList = make([][]string, 0, length)
// 	)
//
// 	for i, row := range sheet.Rows {
// 		if i == 0 {
// 			// 第一行 title 跳过
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
// 		return nil, errors.New("excel行数有误，内容为空")
// 	}
// 	fmt.Printf("excel 行数：%d", len(valuesList))
//
// 	for i, r := range valuesList {
// 		if len(r) < 3 {
// 			fmt.Printf("数据列数不对，excel行：%d,activity id ：%v", i, r)
// 			continue
// 		}
//
// 		var (
// 			itemField = &UpdateLanguageStatusBatchReq{}
// 		)
//
// 		activityId, _ := strconv.ParseInt(r[0], 10, 64)
// 		statusStr := strings.Split(r[2], "-")[0]
// 		status, _ := strconv.ParseInt(statusStr, 10, 64)
//
// 		itemField.ActivityId = activityId
// 		itemField.Language = r[1]_ "klook.libs/logger"

// 		itemField.Status = status
//
// 		importItems = append(importItems, itemField)
// 	}
//
// 	return importItems, nil
// }

func channelTest1(ch chan string, msg string) {
	print("start channelTest1")
	go func() {
		time.Sleep(6*time.Second)
		ch <- msg
		return
	}()
}

func channelTest2(ch chan string) {
	timeOut := time.NewTimer(3 * time.Second).C
	go func() {
		flag := false
		for {
			select {
			case <- timeOut:
				print("\n///////////////////time out //////////////////")
				return
			case msg := <-ch:
				if msg == "done" {
					print("\n work done ")
					flag = true
				} else {
					print( "no done ")
				}
			default:
				print(" wait for msg")
			}

			if flag {
				print("\n<<<<<<<<<<<<<<for over >>>>>>>>>>>>>>>>>\n")
				break
			}
		}
		print("\n all done ")
	}()
}

func main() {
	ch := make(chan string)
	channelTest2(ch)

	channelTest1(ch, "done")
	time.Sleep(3*time.Second)
	print("======== main over ===========")
}
