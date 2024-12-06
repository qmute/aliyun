package testdata

import (
	"os"

	"github.com/qmute/aliyun"
)

var (
	TestAccount         = os.Getenv("Account")
	TestAccessKeyID     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	LogRoleArn          = os.Getenv("LogRoleArn")
)

func Config() aliyun.Config {
	return aliyun.Config{
		Account: TestAccount,
		Key:     TestAccessKeyID,
		Secret:  TestAccessKeySecret,
	}
}
