package main

import "fmt"

func main() {
	// data, err := InitDataCleanValue()
	// if err != nil {
	// 	fmt.Printf("=============== %s", err.Error())
	// }
	//
	// tmpMap := make(map[int64]bool)
	// for _, id := range tmp {
	// 	tmpMap[id] = true
	// }
	//
	// for _, v := range data {
	// 	if _, ok := tmpMap[v.ActivityId]; !ok {
	// 		fmt.Printf("##############")
	// 	}
	// }
	//
	// dt, _ := json.Marshal(data)
	// fmt.Printf("======================>\n %s", string(dt))
	data, err := parseSqlForMfo()
	if err != nil {
		return
	}

	fmt.Printf("======================================")
	langList := []string{"en_US", "en_UK"}
	sql := "insert into mfo_page (keyword_id, city_id,country_id, lang) values (%d, %d, %d, \"%s\");"
	for _, v := range data {
		for _, lang := range langList {
			insertSql := fmt.Sprintf(sql, v.KeywordID, v.CityID, v.CountryId, lang)
			fmt.Println(insertSql)
		}
	}

	// keywprdList := []int64{97, 98, 99, 100, 101, 102, 103, 104, 105}
	// cityIdList := []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 16, 17, 18, 19, 20, 22, 23, 25, 28, 29, 30, 32, 33, 34, 35, 36, 40, 42, 43, 44, 45, 46, 47, 48, 49, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 77, 78, 79, 80, 81, 82, 83, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 103, 106, 107, 108, 109, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 128, 129, 130, 131, 132, 133, 134, 136, 138, 139, 140, 141, 142, 143, 144, 145, 146, 148, 149, 150, 151, 154, 156, 157, 159, 162, 163, 164, 166, 167, 169, 170, 171, 175, 177, 178, 180, 181, 183, 185, 186, 188, 190, 191, 193, 194, 195, 197, 198, 199, 200, 202, 203, 204, 205, 207, 208, 209, 210, 211, 212, 214, 215, 216, 235, 241, 252, 253, 254, 262, 264, 266, 267, 268, 269, 270, 271, 272, 274, 275, 276, 279, 281, 284, 286, 287, 289, 291, 292, 293, 294, 295, 297, 299, 300, 301, 302, 303, 305, 306, 307, 308, 309, 310, 314, 315, 316, 317, 318, 320, 324, 330, 331, 333, 337, 338, 339, 340, 342, 343, 344, 346, 347, 348, 349, 350, 351, 353, 355, 356, 357, 358, 369, 372, 374, 375, 376, 379, 380, 381, 384, 387, 388, 390, 396, 398, 399, 400, 401, 402, 405, 409, 412, 414, 420, 422, 424, 426, 427, 432, 434, 436, 437, 440, 442, 450, 453, 456, 459, 460, 464, 465, 468, 469, 471, 472, 480, 481, 485, 488, 489, 491, 494, 495, 496, 500, 508, 509, 517, 520, 521, 522, 523, 526, 527, 530, 533, 534, 535, 537, 539, 540, 541, 544, 549, 554, 557, 4819, 5066, 6409, 6484, 6488, 6806, 12146, 18056, 23300, 23301, 24144, 24650, 25675, 27456, 29785, 39554, 71287, 76792, 79946, 90932, 110819, 112144, 365089, 365257, 365301, 365305, 365306, 365307, 365309, 365316, 365344, 365346, 365347, 365372, 365389, 365400, 365406, 365463, 700462, 701534, 701739, 10000040, 10000043, 10000044, 10000047}
	// countryIdList := []int64{2, 4, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 31, 32, 33, 34, 35, 37, 38, 39, 41, 42, 43, 44, 47, 48, 49, 50, 52, 53, 55, 56, 58, 59, 60, 61, 62, 63, 66, 67, 68, 69, 70, 72, 75, 84, 85, 86, 91, 92, 93}
	//
	// for _, lang := range langList {
	// 	for _, keyid := range keywprdList {
	// 		for _, city := range cityIdList {
	// 			insSql := fmt.Sprintf("insert into mfo_page (keyword_id, city_id, lang) values (%d, %d, \"%s\");", keyid, city, lang)
	// 			fmt.Println(insSql)
	// 		}
	//
	// 		for _, c := range countryIdList {
	// 			insSql := fmt.Sprintf("insert into mfo_page (keyword_id, country_id, lang) values (%d, %d, \"%s\");", keyid, c, lang)
	// 			fmt.Println(insSql)
	// 		}
	// 	}
	// }
	// respCN := make([]*ReadData, 4000)
	// respHK := make([]*ReadData, 4000)
	// respTW := make([]*ReadData, 4000)
	// err := json.Unmarshal([]byte(dataCn), &respCN)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Print("==========>", len(respCN))
	//
	// err = json.Unmarshal([]byte(dataHk), &respHK)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Print("==========>", len(respHK))
	//
	// err = json.Unmarshal([]byte(dataTw), &respTW)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Print("==========>", len(respTW))
}
