package ram_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/qmute/aliyun/ram"
	"github.com/qmute/aliyun/testdata"
)

var _ = Describe("Ram", func() {

	Context("AssumeRole", func() {
		It("200", func() {
			opts := []ram.AssumeRoleOpt{
				ram.WithAssumeRoleArn(testdata.LogRoleArn),
				ram.WithAssumeRoleSessionName("test"),
				ram.WithAssumeRoleDurationSeconds(1800),
				ram.WithAssumeRolePolicy(`
{
  "Version": "1",
  "Statement": [
    {
      "Action": "log:PostLogStoreLogs",
      "Resource": "acs:log::1546067273439950:project/qian-app/logstore/qian-app-log/",
      "Effect": "Allow"
    }
  ]
}

`),
			}
			rsp, err := clt.AssumeRole(opts...)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rsp.RequestId).ShouldNot(BeEmpty())
			Expect(rsp.AssumedRoleUser.AssumedRoleId).ShouldNot(BeEmpty())
			Expect(rsp.AssumedRoleUser.Arn).ShouldNot(BeEmpty())
			Expect(rsp.Credentials.AccessKeyId).ShouldNot(BeEmpty())
			Expect(rsp.Credentials.AccessKeySecret).ShouldNot(BeEmpty())
			Expect(rsp.Credentials.SecurityToken).ShouldNot(BeEmpty())
			Expect(rsp.Credentials.Expiration).ShouldNot(BeEmpty())
		})

		Context("400", func() {
			hasErr := func(suberr string, opt ...ram.AssumeRoleOpt) {
				Expect(opt).ShouldNot(HaveLen(0))
				rsp, err := clt.AssumeRole(opt...)
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring(suberr))
				Expect(rsp).Should(BeEquivalentTo(ram.AssumeRoleResponse{}))
			}

			It("roleArn", func() {
				hasErr("roleArn must be not empty", ram.WithAssumeRoleArn(""))
			})

			It("roleSessionName", func() {
				hasErr("roleSessionName must be not empty",
					ram.WithAssumeRoleSessionName(""))
			})

			It("durationSeconds", func() {
				hasErr("过期时间最小值为900秒",
					ram.WithAssumeRoleArn("x"),
					ram.WithAssumeRoleSessionName("y"),
					ram.WithAssumeRoleDurationSeconds(899))
			})

		})
	})
})
