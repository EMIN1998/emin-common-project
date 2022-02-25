package main

import "fmt"

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

type ContactConfig struct {
	FlagCountry     bool `json:"flag_country" db:"flag_country"`
	FlagEnglish     bool `json:"flag_english" db:"flag_english"`
	FlagMailbox     bool `json:"flag_mailbox" db:"flag_mailbox"`
	FlagLastname    bool `json:"flag_lastname" db:"flag_lastname"`
	FlagFirstname   bool `json:"flag_firstname" db:"flag_firstname"`
	FlagCountryCode bool `json:"flag_country_code" db:"flag_country_code"`
	FlagPhoneNumber bool `json:"flag_phone_number" db:"flag_phone_number"`
}

type CategoryContactConfig struct {
	CategoryId int64 `json:"category_id" db:"category_id"`
	ContactConfig
}

// func testfunc() {
// 	var str string
// 	str = "{\n    \"category_id\": 15,\n    \"flag_country\": true,\n    \"flag_english\": true,\n    \"flag_mailbox\": true,\n    \"flag_lastname\": true,\n    \"flag_firstname\": true,\n    \"flag_country_code\": true,\n    \"flag_phone_number\": true\n}"
//
// 	var tmp CategoryContactConfig
// 	err := json.Unmarshal([]byte(str), &tmp)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
//
// 	fmt.Print(tmp)
// }
// func ToIdString(ids []int64) string {
// 	stList := make([]string, 0)
// 	for _, id := range ids {
// 		st := fmt.Sprintf("%d", id)
// 		stList = append(stList, st)
// 	}
//
// 	str := strings.Join(stList, ",")
// 	return str
// }

func InInt64Slice(item int64, slice []int64) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}
	return false
}

var allSubCategoryId = []int64{1, 2, 4, 7, 8, 9, 14, 15, 16, 17, 18, 19, 20, 21, 22, 168, 171, 177, 179, 185, 187, 192, 194, 198, 206, 211, 220, 223, 228, 230, 232, 249, 251, 254, 256, 258, 260, 263, 266, 275, 279, 282, 287, 298, 300, 303, 306, 309, 311, 313, 317, 319, 321, 323, 325, 329, 331, 333, 337, 341, 343, 345, 353, 356, 358, 360, 362, 364, 366, 369, 372, 374, 376, 378, 380, 382, 384, 386, 388, 390, 392, 394, 396, 398, 400, 402, 404, 406, 409, 410, 413, 415, 426, 439, 452, 523, 525, 527, 529, 531, 534, 535, 540, 542, 551, 552, 558}

var type1 = []int64{1, 2, 4, 7, 8, 9, 14, 15, 16, 20, 21, 22, 187, 192, 206, 211, 220, 223, 228, 230, 232, 249, 251, 254, 256, 258, 260, 263, 266, 275, 279, 282, 287, 300, 303, 306, 309, 311, 313, 331, 356, 360}

var type2 = []int64{1, 2, 4, 7, 8, 9, 14, 15, 16, 17, 18, 19, 20, 21, 22, 168, 171, 177, 179, 185, 187, 192, 194, 198, 206, 211, 220, 223, 228, 230, 232, 249, 251, 254, 256, 258, 260, 263, 266, 275, 279, 282, 287, 298, 300, 303, 306, 309, 311, 313, 317, 319, 321, 323, 325, 329, 331, 333, 341, 343, 345, 353, 356, 358, 360, 364}

var type3 = []int64{1, 2, 4, 7, 8, 9, 14, 15, 16, 17, 18, 19, 20, 21, 22, 168, 171, 177, 179, 187, 194, 206, 220, 223, 228, 230, 232, 251, 263, 266, 279, 282, 287, 298, 300, 303, 306, 309, 311, 313, 317, 319, 321, 323, 325, 331, 413, 415, 523}

func main() {
	target := make([]int64, 0)
	for _, v := range allSubCategoryId {
		if !InInt64Slice(v, type3) {
			if InInt64Slice(v, type1) && InInt64Slice(v, type2) {
				target = append(target, v)
			}
		}

		if !InInt64Slice(v, type2) && !InInt64Slice(v, type1) && !InInt64Slice(v, type3) {
			target = append(target, v)
		}
	}

	fmt.Print(target)
}
