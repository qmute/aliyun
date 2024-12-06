package sms_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun"
	"github.com/qmute/aliyun/sms"
)

var (
	testAccessKeyID     = os.Getenv("AccessKeyId")
	testAccessKeySecret = os.Getenv("AccessKeySecret")
	testMobile          = os.Getenv("Mobile")
	testSignName        = os.Getenv("SignName")
)

func TestSms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sms Suite")
}

var (
	clt *sms.Client
	ctx context.Context
)

var _ = BeforeSuite(func() {
	ctx = context.Background()
	conf := aliyun.Config{
		Account: "",
		Key:     testAccessKeyID,
		Secret:  testAccessKeySecret,
	}

	clt = sms.New(conf)
})
