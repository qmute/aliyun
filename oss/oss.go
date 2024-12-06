package oss

import (
	"github.com/denverdino/aliyungo/oss"

	"github.com/qmute/aliyun"
)

const (
	defaultExpire = 30
)

type Option struct {
	Bucket   string
	Dir      string
	Endpoint string
	Region   string
	Base     string // oss 文件下载地址的基地址
	Expire   int64  // 过期时间, 秒
	Internal bool   // 是否内网
}

func (p *Option) Valid() error {
	if p.Bucket == "" {
		return ErrBucketEmpty
	}

	if p.Endpoint == "" {
		return ErrEndpointEmpty
	}

	if p.Region == "" {
		return ErrRegionEmpty
	}

	return nil
}

// New 新建一个oss 客户端
func New(conf aliyun.Config, opt ...OptionFunc) *Client {
	if err := conf.Valid(); err != nil {
		panic(err)
	}

	clt := &Client{
		account: conf.Account,
		key:     conf.Key,
		secret:  conf.Secret,
		opt: Option{
			Dir:    "/",
			Expire: defaultExpire,
		},
	}

	clt.init(opt...)

	return clt
}

type Client struct {
	account string
	key     string
	secret  string

	opt Option

	bucket *oss.Bucket // 真实操作bucket

}

func (p *Client) Valid() error {
	if err := p.opt.Valid(); err != nil {
		return err
	}

	return nil
}

func (p *Client) Opt() Option {
	return p.opt
}

func (p *Client) Key() string {
	return p.key
}

func (p *Client) Secret() string {
	return p.secret
}

func (p *Client) BucketClient() *oss.Bucket {
	return p.bucket
}

// R 创建一个新的客户端实例
func (p *Client) R(opt ...OptionFunc) *Client {
	clt := &Client{
		account: p.account,
		key:     p.key,
		secret:  p.secret,
		opt:     p.opt,
	}

	clt.init(opt...)
	return clt
}

func (p *Client) init(opt ...OptionFunc) *Client {
	for _, o := range opt {
		o(p)
	}

	if err := p.Valid(); err != nil {
		panic(err)
	}

	clt := oss.NewOSSClient(oss.Region(p.opt.Region), p.opt.Internal,
		p.key, p.secret, true)
	p.bucket = clt.Bucket(p.opt.Bucket)

	return p
}
