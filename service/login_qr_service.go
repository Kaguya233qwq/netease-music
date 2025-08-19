package service

import (
	"encoding/json"

	"github.com/go-musicfox/netease-music/util"
)

type LoginQRService struct {
	UniKey string `json:"unikey"`
}

func (service *LoginQRService) GetKey() (float64, []byte, string) {
	data := map[string]string{
		"type": "1",
	}

	api := "https://music.163.com/weapi/login/qrcode/unikey"
	code, bodyBytes := util.CallWeapi(api, data)
	if code != 200 || len(bodyBytes) == 0 {
		return code, bodyBytes, ""
	}
	_ = json.Unmarshal(bodyBytes, service)

	// 生成 chainId，这个是新版本新加的参数
	cookieJar := util.GetGlobalCookieJar()
	chainID := util.GenerateChainID(cookieJar)
	qrcodeUrl := ("http://music.163.com/login?codekey=" +
		service.UniKey + "&chainId=" + chainID)
	return code, bodyBytes, qrcodeUrl
}

func (service *LoginQRService) CheckQR() (float64, []byte) {
	if service.UniKey == "" {
		return 0, nil
	}
	data := map[string]string{
		"type": "1",
		"key":  service.UniKey,
	}

	api := "https://music.163.com/weapi/login/qrcode/client/login"
	code, bodyBytes := util.CallWeapi(api, data)
	return code, bodyBytes
}
