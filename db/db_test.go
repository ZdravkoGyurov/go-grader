package db_test

import (
	"testing"

	"grader/app"
	"grader/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var _ = Describe("DB", func() {
	var (
		appCtx app.Context
	)

	BeforeSuite(func() {
		appCtx = app.NewContext()
	})

	Describe("Connecting to mongo DB", func() {
		It("Should connect and ping successfully", func() {
			_, err := db.Connect(appCtx)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
