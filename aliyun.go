package aliyun

import "errors"

var (
	ErrAccountEmpty = errors.New("account must be not empty")
	ErrKeyEmpty     = errors.New("key must be not empty")
	ErrSecretEmpty  = errors.New("secret must be not empty")
)

type Config struct {
	Account string
	Key     string
	Secret  string
}

func (p *Config) Valid() error {
	// if p.Account == "" {
	// 	return ErrAccountEmpty
	// }

	if p.Key == "" {
		return ErrKeyEmpty
	}

	if p.Secret == "" {
		return ErrSecretEmpty
	}

	return nil
}
