package sms

import (
	"fmt"
	"strings"

	"github.com/denverdino/aliyungo/sms"
)

const (
	Success     = "OK"
	numberLimit = 1000
)

type SendResult sms.SendSmsResponse

// type SendResponse struct {
// 	sms.SendSmsResponse
// }

func (p *SendResult) IsOk() bool {
	return p.Code == Success
}

func (p *SendResult) Error() string {
	return fmt.Sprintf("send sms fail: %s, RequestId<%s>, Code<%s>, BizId<%s>",
		p.Message, p.RequestId, p.Code, p.BizId)
}

type SendArgs struct {
	// 必填:待发送手机号
	// 批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
	PhoneNumbers []string
	// 必填
	SignName string
	// 必填:短信模板-可在短信控制台中找到
	TemplateCode string
	// 可选:模板中的变量替换JSON串,如模板内容为"亲爱的${name},您的验证码为${code}"时,此处的值为
	// 友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\\r\\n,
	// 否则会导致JSON在服务端解析失败
	TemplateParam string
	// 可选-上行短信扩展码(扩展码字段控制在7位或以下，无特殊需求用户请忽略此字段)
	SmsUpExtendCode string
	// 可选:outId为提供给业务方扩展字段,最终在短信回执消息中将此值带回给调用者
	OutId string
}

func (p *SendArgs) Valid() error {
	size := len(p.PhoneNumbers)
	if size == 0 {
		return ErrEmptyPhoneNumbers
	}
	if size > numberLimit {
		return ErrLimitPhoneNumbers
	}

	if len(strings.TrimSpace(p.SignName)) == 0 {
		return ErrEmptySignName
	}

	if len(strings.TrimSpace(p.TemplateCode)) == 0 {
		return ErrEmptyTemplateCode
	}

	return nil
}

func (p *SendArgs) ToArgs() *sms.SendSmsArgs {
	return &sms.SendSmsArgs{
		PhoneNumbers:    strings.Join(p.PhoneNumbers, ","),
		SignName:        p.SignName,
		TemplateCode:    p.TemplateCode,
		TemplateParam:   p.TemplateParam,
		SmsUpExtendCode: p.SmsUpExtendCode,
		OutId:           p.OutId,
	}
}
