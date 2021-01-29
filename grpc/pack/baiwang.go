package pack

import (
	"crypto/md5"
	"encoding/json"
	pb "fapiao/Ys_Pb_Fapiao"
	"fapiao/tools"
	"fmt"
	"github.com/tiaguinho/gosoap"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const jsonUrl string = "https://dev.fapiao.com:18944/fpt-dsqz/services/DZFPJsonService?wsdl"

var appID string = "6d29f136114544bcc73edcce960c430231183cc192c433e2b9ebcad56e8ceb08"
var xsfNsrsbh string = "110109500321655"
var xsfMc string = "百旺电子测试2"
var contentPassword = "5EE6C2C11DD421F2"

var requestCode = "DZFPQZ"

type makeInvoiceResult struct {
	FPQQLSH string `json:"FPQQLSH"`
	FP_DM   string `json:"FP_DM"`
	FP_HM   string `json:"FP_HM"`
	KPRQ    string `json:"KPRQ"`
	JYM     string `json:"JYM"`
	PDF_URL string `json:"PDF_URL"`
	SP_URL  string `json:"SP_URL"`
}

type makeInvoiceResponse struct {
	makeInvoiceResult string
}

//生成请求号
func makeInvoiceTransCode() string {
	str := "WW" + strconv.Itoa(tools.GetNowIntDate(0)) + strconv.Itoa(rand.Intn(8998)+1001)
	return str
}

func makeContentKey(base64Content string, password string) string {
	data := []byte(base64Content)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	aesStr := tools.AESEncrypt([]byte(md5Str), []byte(password))
	return string(aesStr)
}

// 开票
func MakeInvoiceThirdAPI(aP pb.Apply, iM pb.InvoiceModel) error {
	interFaceCode := "DFXJ1001"

	transCode := makeInvoiceTransCode()

	requestTime := tools.GetNowStringDate()

	dataExchangeID := requestCode + interFaceCode + strconv.Itoa(tools.GetNowIntDate(1)) + tools.GetRandomString(9)

	generateData := generateData()

	content1 := `
			{
			   "REQUEST_COMMON_FPKJ": {
					"FPQQLSH": "%s", 
					"BMB_BBH": "",
					"ZSFS": "0",
					"KPLX": "0",
					"XSF_NSRSBH": "%s",
					"XSF_MC": "%s",
					"XSF_DZDH": "北京市朝阳区望京东园四区浦项中心A座20层、010-55993423",
					"XSF_YHZH": "",
					"GMF_NSRSBH": "%s",
					"GMF_MC": "%s",
					"GMF_DZDH": "%s",
					"GMF_YHZH": "%s",
					"GMF_SJH": "",
					"GMF_DZYX": "%s",
					"FPT_ZH": "",
					"WX_OPENID": "",
					"KPR": "申艳平",
					"SKR": "雷蓓玉",
					"FHR": "牛亚男",
					"YFP_DM": "",
					"YFP_HM": "",
					"JSHJ": "%d",
					"HJJE": "%d",
					"HJSE": "%d",
					"KCE": "",
					"BZ": "%s",
					"HYLX": "",
					"BY1": "备用字段 1",
					"BY2": "备用字段 2",
					"BY3": "备用字段 3",
					"BY4": "备用字段 4",
					"BY5": "备用字段 5",
					"BY6": "备用字段 6",
					"BY7": "备用字段 7",
					"BY8": "备用字段 8",
					"BY9": "备用字段 9",
					"BY10": "备用字段 10",
					"WX_ORDER_ID": "",
					"WX_APP_ID": "", 
					"ZFB_UID":"",
					"TSPZ": "00",
					"QJ_ORDER_ID": "",
					"WX_GROUP_ID": "",
					"COMMON_FPKJ_XMXXS": {
						"COMMON_FPKJ_XMXX": {
							"FPHXZ": "0", 
							"SPBM": "",
							"ZXBM": "", 
							"YHZCBS": "0", 
							"LSLBS": "1", 
							"ZZSTSGL": "",
							"XMMC": "%s", 
							"GGXH": "", 
							"DW": "", 
							"XMSL": "%d", 
							"XMDJ": "%d", 
							"XMJE": "%d", 
							"SL": "%d", 
							"SE": "%d", 
							"BY1": "备用字段 1", 
							"BY2": "备用字段 2", 
							"BY3": "备用字段 3", 
							"BY4": "备用字段 4", 
							"BY5": "备用字段 5"
						} 
					}
				}
			}
	`
	content := fmt.Sprintf(content1,
		transCode,
		xsfNsrsbh,
		xsfMc,
		aP.InvoiceCode,
		aP.InvoiceHead,
		aP.CompanyAddress,
		aP.CompanyBankNo,
		aP.UserEmail,
		iM.InvoiceMoneyNoTax,
		iM.InvoiceMoney,
		iM.InvoiceTax,
		aP.Remark,
		InvoiceContentDesc[iM.InvoiceContentType],
		iM.ProjectNum,
		iM.ProjectUnitPrice,
		iM.InvoiceMoneyNoTax,
		iM.InvoiceTaxRate,
		iM.InvoiceTax,
	)

	base64Content := tools.Base64EncodeString(content)

	contentKey := makeContentKey(base64Content, contentPassword)
	contentKey = tools.Base64EncodeString(contentKey)
	fmt.Println("contentKey:", contentKey)

	generateData = fmt.Sprintf(generateData, appID, interFaceCode, requestCode, requestTime, dataExchangeID, base64Content, contentKey)
	fmt.Println("generateData:", generateData)
	err := requestThirdApi(generateData)
	if err != nil {
		return err
	}
	return nil
}

// 红冲接口
func TurnRedInvoiceThirdAPI(iM pb.InvoiceModel) error {
	//interFaceCode := "DFXJ1005"
	return nil
}

// 全局请求信息
func generateData() string {
	str := `
		{
			"interface": {
				"globalInfo": {
					"appId": "%s",
					"interfaceId": "",
					"interfaceCode": "%s",
					"requestCode": "%s",
					"requestTime": "%s",
					"responseCode": "DS",
					"dataExchangeId": "%s"
				},
				"returnStateInfo": {
					"returnCode": "",
					"returnMessage": ""
				},
				"Data": {
					"dataDescription": {
						"zipCode": "0"
					},
					"content": "%s",
					"contentKey": "%s"
				}
			}
		}
	`
	return str
}

func requestThirdApi(generateData string) error {
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(jsonUrl, httpClient)
	if err != nil {
		return err
	}

	params := gosoap.Params{
		"local_cert": "../testISSUE.pem",
		"passphrase": "123456",
		"in0":        generateData,
	}
	res, err := soap.Call("doJsonService", params)

	fmt.Println("res:", res, err)

	if err != nil {
		return err
	}

	var r makeInvoiceResponse

	err = res.Unmarshal(&r)
	if err != nil {
		return err
	}

	result := makeInvoiceResult{}
	err = json.Unmarshal([]byte(r.makeInvoiceResult), &result)
	fmt.Println("call response:", result)
	return nil
}
