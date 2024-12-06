package ram_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun/ram"
	"github.com/qmute/aliyun/testdata"
)

func TestRam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ram Suite")
}

var clt *ram.Client
var _ = BeforeSuite(func() {
	conf := testdata.Config()
	clt = ram.New(conf)
})
