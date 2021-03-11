package db_test

import (
	"context"
	"grader/db"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var _ = Describe("DB", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	BeforeSuite(func() {
		ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	})

	AfterSuite(func() {
		cancel()
	})

	Describe("Connecting to mongo DB", func() {
		It("Should connect and ping successfully", func() {
			_, err := db.Connect(ctx)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
