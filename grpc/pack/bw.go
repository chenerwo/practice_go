package main

import (
"fmt"
	//curl "github.com/andelf/go-curl"
	"encoding/json"
)

type requestBaiwangResult struct {
	Interface requestBaiwangInterFace `json:"interface"`
}

type requestBaiwangInterFace struct {
	globalInfo struct {
		InterfaceCode  string `json:"interfaceCode"`
		ResponseCode   string `json:"responseCode"`
		AppId          string `json:"appId"`
		RequestTime    string `json:"requestTime"`
		RequestCode    string `json:"requestCode"`
		InterfaceId    string `json:"interfaceId"`
		DataExchangeId string `json:"dataExchangeId"`
	}
	returnStateInfo struct {
		ReturnCode    string `json:"returnCode"`
		ReturnMessage string `json:"returnMessage"`
	}
	Data struct {
		dataDescription struct {
			zipCode string
		}
		Content string `json:"content"`
	}
}

func main() {
	var b
	fmt.Println(b/100)
	//b = fmt.Sprintf("num:%.2f", )

	//fmt.Println(b)
	/*
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://www.baidu.com/")

	// make a callback function
	fooTest := func (buf []byte, userdata interface{}) bool {
		println("DEBUG: size=>", len(buf))
		//println("DEBUG: content=>", string(buf))
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	 */

	json := `{"interface":{"globalInfo":{"interfaceCode":"DFXJ1001","responseCode":"1","appId":"0","requestTime":"2020-07-16 10:11:21:415","requestCode":"DZFPQZ","interfaceId":"","dataExchangeId":"DZFPQZDFXJ10012020716byna55mz3"},"returnStateInfo":{"returnCode":"0000","returnMessage":"成功"},"Data":{"dataDescription":{"zipCode":"0"},"content":"eyJGUFFRTFNIIjoiV1cyMDIwMDcxNjEwMTEyMDY3MDAiLCJGUF9ETSI6IjE1MjAwMDE4NjM1NyIsIkZQX0hNIjoiMzA0NTA2MTQiLCJLUFJRIjoiMjAyMDA3MTYxMDAxNTUiLCJKWU0iOiIwODczNTI1NjgwOTU1MzQ2MDM1NCIsIlBERl9VUkwiOiJoRwOi8vZGV2LmZhcGlhby5jb206MTkwODAvZHpmcC13ZWIvcGRmL2Rvd25sb2FkP3JlcXVlc3Q9ZTV1aGY4V0VUSU9NZ2FhMmNDVU10cExTZTR4SXQ3T3d2MXJRWjY0MjM0TUpUcUtxdTdwWVZGOFVZVDNacDdPaVAxb1FnMVRHVUNDMUp3eHR5OFlzY3dfXyU1RWVCaWJFZmFlR2IiLCJTUF9VUkwiOiJodHRwczovL3Rlc3R3eC5mYXBpYW8uY29tL2ZwdC13ZWNoYXQvd3hhZGRjYXJkLmRvP2NvZGU9U1FsVXhycE1LSmNyJTJCRFFRR2pyUWcyNERkdXg2NGNMU2dLTXB3Yk8lMkZYZjkzck1IcVRiS1RoMFZ5SWVSYWYxSVVqM0dOaWxrMDZHQkklMEFnTVNkem80TkJBJTNEJTNEIn0="}}}`

	var result requestBaiwangResult
	err := json
}
