package main

import (
	"fmt"
	"time"
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
}
