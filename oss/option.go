package oss

// OptionFunc 是一个参数选项 interface.
type OptionFunc func(clt *Client)

func WithBucket(bucket string) OptionFunc {
	return func(clt *Client) {
		if bucket == "" {
			panic("Bucket must be not empty")
		}

		clt.opt.Bucket = bucket
	}
}

func WithDir(dir string) OptionFunc {
	return func(clt *Client) {
		if dir == "" {
			panic("Dir must be not empty")
		}

		clt.opt.Dir = dir
	}
}

func WithEndpoint(endpoint string) OptionFunc {
	return func(clt *Client) {
		clt.opt.Endpoint = endpoint
	}
}

func WithRegion(region string) OptionFunc {
	return func(clt *Client) {
		clt.opt.Region = region
	}
}

func WithInternal(internal bool) OptionFunc {
	return func(clt *Client) {
		clt.opt.Internal = internal
	}
}

func WithBase(base string) OptionFunc {
	return func(clt *Client) {
		clt.opt.Base = base
	}
}

func WithExpire(expire int64) OptionFunc {
	return func(clt *Client) {
		clt.opt.Expire = expire
	}
}
