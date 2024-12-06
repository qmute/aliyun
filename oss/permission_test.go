package oss_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Permission", func() {

	It("WEB", func() {
		token, err := clt.WebToken()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(token).ShouldNot(BeNil())
		Expect(token.BaseUrl).ShouldNot(BeEmpty())
		Expect(token.PolicyToken).ShouldNot(BeNil())
	})

	It("STS", func() {
		token, err := clt.StsToken()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(token.BaseUrl).ShouldNot(BeEmpty())
		Expect(token.Credentials).ShouldNot(BeNil())
	})
})
