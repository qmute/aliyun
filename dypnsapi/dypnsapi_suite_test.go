package dypnsapi_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun"
	"github.com/qmute/aliyun/dypnsapi"
)

var (
	testAccessKeyID     = os.Getenv("AccessKeyId")
	testAccessKeySecret = os.Getenv("AccessKeySecret")
	testMobile          = os.Getenv("Mobile")
)

func TestDypnsapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dypnsapi Suite")
}

var (
	clt *dypnsapi.Client
)
var _ = BeforeSuite(func() {
	conf := aliyun.Config{
		Account: "",
		Key:     testAccessKeyID,
		Secret:  testAccessKeySecret,
	}

	clt = dypnsapi.New(conf)
})
