package util

import (
	"fmt"
	"strconv"
	"strings"
)

// 月份转换Map
var monthCorrespond = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

// ToIso8601 时间戳转换
//
//	@param timeList
//	@return string
//	@author KevinMatt 2021-07-25 13:42:29
//	@function_mark PASS
func ToIso8601(timeList []string) string {
	temp := fmt.Sprintf("%s-%s-%sT%s+%s:%s", timeList[3], monthCorrespond[timeList[0]], timeList[1], timeList[2], timeList[4][1:3], timeList[4][3:])
	return temp
}

// ConCatStrings 字符串高效拼接
//
//	@param stringList ...string
//	@return string
//	@author: Kevineluo 2022-07-31 12:52:54
func ConCatStrings(stringList ...string) string {
	var builder strings.Builder
	for index := range stringList {
		builder.WriteString(stringList[index])
	}
	return builder.String()
}

// QuatToNum
//
//	@param text string
//	@return sum int
//	@author: Kevineluo 2022-07-31 12:24:09
func QuatToNum(text string) (sum int) {
	for index := 0; index < len(text); index++ {
		if text[index] == ',' {
			continue
		}
		temp, _ := strconv.Atoi(string(text[index]))
		sum = sum*10 + temp
	}
	return
}
