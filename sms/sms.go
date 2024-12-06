package sms

import (
	"context"

	"github.com/denverdino/aliyungo/common"
	"github.com/denverdino/aliyungo/sms"

	"github.com/qmute/aliyun"
)

// New return a new sms client
func New(conf aliyun.Config) *Client {
	if err := conf.Valid(); err != nil {
		panic(err)
	}

	dysmsClt := sms.NewDYSmsClient(conf.Key, conf.Secret)
	dysmsClt.Region = common.Beijing

	return &Client{
		dySmsClient: dysmsClt,
	}
}

// Client sms client
type Client struct {
	dySmsClient *sms.DYSmsClient
}

func (p *Client) Send(ctx context.Context, args *SendArgs) (*SendResult, error) {
	if err := args.Valid(); err != nil {
		return nil, err
	}

	rsp, err := p.dySmsClient.SendSms(args.ToArgs())
	if err != nil {
		return nil, err
	}

	ret := SendResult(*rsp)
	if ret.IsOk() {
		return &ret, nil
	}

	return nil, &ret
}
