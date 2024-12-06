package sms_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun/sms"
)

var _ = Describe("Sms", func() {
	Context("SendSms", func() {
		var args *sms.SendArgs
		BeforeEach(func() {
			args = &sms.SendArgs{
				PhoneNumbers:    []string{testMobile},
				SignName:        testSignName,
				TemplateCode:    "SMS_122289939",
				TemplateParam:   "{\"code\": \"boOob\"}",
				SmsUpExtendCode: "",
				OutId:           "",
			}
		})

		It("200", func() {
			ret, err := clt.Send(ctx, args)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(ret).ShouldNot(BeNil())
			Expect(ret.Code).Should(Equal(sms.Success))
		})

		It("400", func() {
			args.SignName = ""
			ret, err := clt.Send(ctx, args)
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).Should(Equal(sms.ErrEmptySignName.Error()))
			Expect(ret).Should(BeNil())
		})
	})
})
