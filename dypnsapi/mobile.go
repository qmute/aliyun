package dypnsapi

import (
	"log"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dypnsapi"
	"github.com/pkg/errors"
)

const (
	ResponseOk = "OK"
	VerifyPass = "PASS"
)

const (
	// VerifyResultPass 一致
	VerifyResultPass = "PASS"
	// VerifyResultReject 不一致
	VerifyResultReject = "REJECT"
	// VerifyResultUnknown 无法判断
	VerifyResultUnknown = "UNKNOWN"
)

// VerifyResult 验证结果
type VerifyResult struct {
	Result string
}

func (p VerifyResult) Pass() bool {
	return p.Result == VerifyResultPass
}

// VerifyMobile 验证是否本机手机号
// 1. accessCode 只能使用一次
// 2. accessCode 第二次使用时，返回验证结果为UNKNOWN
// 3. accessCode 必须与手机号相区配，否则返回错误
// 4. isp.SYSTEM_ERROR 可能是accessCode 格式错误
func (p *Client) VerifyMobile(accessCode string, phoneNumber string) (VerifyResult, error) {
	result := VerifyResult{}
	req := dypnsapi.CreateVerifyMobileRequest()
	req.PhoneNumber = phoneNumber
	req.AccessCode = accessCode
	req.SetContentType(p.format)
	req.Scheme = p.scheme

	rsp, err := p.dyClt.VerifyMobile(req)
	if err != nil {
		return result, err
	}

	log.Printf("rsp: %#v", rsp)
	if strings.ToUpper(rsp.Code) != ResponseOk {
		return result, errors.Errorf("%s:%s:%s", rsp.RequestId, rsp.Code, rsp.Message)
	}

	result.Result = rsp.GateVerifyResultDTO.VerifyResult

	return result, nil

}
