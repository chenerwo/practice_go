package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	arr := []string{"123", "234", "345"}
	fmt.Println(JoinStringArr2String(arr, ","))
	fmt.Println(String2Timestamp("2020-03-12 12:20:01"))
	s := []int32{1, 2, 3}
	str := []string{"a", "b", "c"}
	fmt.Println("join:", strings.Join(str, ","))
	fmt.Println("inSlice:", valIsInSlice(s, 2))
	fmt.Println("longStr:", strlong())

	fmt.Println(fmt.Sprintf("%.2f", 15999/1000))
}

// 将字符串数组转换成指定分隔符的字符串
func JoinStringArr2String(arr []string, sep string) string {
	arrLen := len(arr)
	if arrLen <= 0 {
		return ""
	}
	var result string
	for k, value := range arr {
		if k > 0 {
			result = fmt.Sprintf("%s%s'%s'", result, sep, value)
		} else {
			result = fmt.Sprintf("'%s'", value)
		}
	}

	return result
}

func String2Timestamp(toBeCharge string) int32 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
	sr := theTime.Unix()
	return int32(sr)
}

func valIsInSlice(s []int32, val int32) bool {
	l := len(s)
	if l <= 0 {
		return false
	}
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func strlong() string {
	str := `
		{
			"interface": {
				"globalInfo": {
					"appId": "百望电子提供 appId",
					"interfaceId": "",
					"interfaceCode": "DFXJ1001",
					"requestCode": "DZFPQZ",
					"requestTime": "2013-09-04 09:58:53",
					"responseCode": "DS",
					"dataExchangeId": "DZFPQZDFXJ10012017033098A6123D0"
				},
				"returnStateInfo": {
					"returnCode": "0000(成功) 0001-9999(错误码)",
					"returnMessage": "成功或错误的详细消息"
				},
				"Data": {
					"dataDescription": {
						"zipCode": "0"
					},
					"content": "Base64 返回数据内容",
					"contentKey": "请求数据 content 节点内的内容 MD5 再 AES 加密"
				}
			}
		}
	`
	return str
}