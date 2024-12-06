package dypnsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dypnsapi"

	"github.com/qmute/aliyun"
)

type Client struct {
	key    string
	secret string
	region string
	format string // 返回值的类型，支持JSON与XML，默认为JSON
	scheme string // 默认https

	dyClt *dypnsapi.Client
}

func New(conf aliyun.Config, opt ...Option) *Client {
	if err := conf.Valid(); err != nil {
		panic(err)
	}

	clt := &Client{
		key:    conf.Key,
		secret: conf.Secret,
		region: "",
		format: "JSON",
		scheme: "https",
	}

	dyClt, err := dypnsapi.NewClientWithAccessKey(clt.region, clt.key, clt.secret)
	if err != nil {
		panic(err)
	}

	clt.dyClt = dyClt

	return clt
}

func (p *Client) init(opt ...Option) *Client {
	for _, o := range opt {
		o(p)
	}

	return p
}
