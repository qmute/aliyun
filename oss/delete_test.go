package oss_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun/oss"
)

var _ = Describe("Delete", func() {

	var info *oss.UploadInfo
	BeforeEach(func() {

		b := []byte(`a\nb\nc\n`)
		info = &oss.UploadInfo{
			Payload:      bytes.NewReader(b),
			OriginalName: "xx111111.txt",
			ContentType:  "",
			Dir:          "",
			Size:         int64(len(b)),
		}

	})

	It("200 - Dir empty", func() {
		ret, err := clt.Upload(ctx, info)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(ret).ShouldNot(BeNil())
		Expect(ret.DownloadUrl).ShouldNot(BeEmpty())
		Expect(ret.Filename).ShouldNot(BeEmpty())
		Expect(ret.Size).Should(Equal(info.Size))

		err = clt.Delete(ctx, ret.DownloadUrl)
		Expect(err).ShouldNot(HaveOccurred())
	})
})
