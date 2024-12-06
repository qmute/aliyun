package oss_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun"
	"github.com/qmute/aliyun/oss"
)

var (
	TestAccount         = os.Getenv("Account")
	TestAccessKeyID     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
)

var (
	clt *oss.Client
	ctx context.Context
)

func init() {
	ctx = context.Background()
	conf := aliyun.Config{
		Account: TestAccount,
		Key:     TestAccessKeyID,
		Secret:  TestAccessKeySecret,
	}

	opts := []oss.OptionFunc{
		oss.WithBase("https://f.weiliplus.com"),
		oss.WithBucket("weiliplus"),
		oss.WithDir("osssdk/"),
		oss.WithRegion("oss-cn-beijing"),
		oss.WithEndpoint("oss-cn-beijing.aliyuncs.com"),
		oss.WithExpire(60),
	}

	clt = oss.New(conf, opts...)
}

func TestOss(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Oss Suite")
}
