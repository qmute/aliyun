package oss_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun/oss"
)

var _ = Describe("Object", func() {
	It("200", func() {
		u := "https://cdn.qian.fm/t/upload/manage/20210820/3a7c040e-019b-11ec-8aee-00163e106f77.jpeg"
		info, err := oss.GetImageInfo(u)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(info.FileSize.Value).Should(Equal("100439"))
		Expect(info.GetFileSize()).Should(Equal(100439))
		Expect(info.Format.Value).Should(Equal("jpg"))
		Expect(info.ImageWidth.Value).Should(Equal("550"))
		Expect(info.GetImageWidth()).Should(Equal(550))
		Expect(info.ImageHeight.Value).Should(Equal("800"))
		Expect(info.GetImageHeight()).Should(Equal(800))
	})
})
