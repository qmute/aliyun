package oss

import (
	"context"
	"strings"
)

type DeleteInfo struct {
	Url []string
}

func (p *Client) Delete(ctx context.Context, url string) error {
	if url == "" {
		return ErrEmptyDelUrl
	}
	if !strings.HasPrefix(url, p.opt.Base) {
		return ErrNotSameBaseUrl
	}

	path := strings.Replace(url, p.opt.Base, "", -1)

	err := p.bucket.Del(path)
	if err != nil {
		return err
	}

	return nil
}
