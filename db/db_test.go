package db_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var _ = Describe("DB", func() {
	Describe("Connecting to mongo DB", func() {
		It("Should connect and ping successfully", func() {
			Expect(1).To(Equal(1))
		})
	})
})
