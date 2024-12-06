package dypnsapi

import "strings"

// Option 是一个参数选项 interface.
type Option func(clt *Client)

func WithRegion(region string) Option {
	return func(clt *Client) {
		clt.region = region
	}
}

func WithFormat(format string) Option {
	return func(clt *Client) {
		clt.format = strings.ToUpper(format)
	}
}
