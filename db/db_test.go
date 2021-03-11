package db_test

import (
	"testing"
	"time"

	"grader/app/config"
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
		cfg config.Config
	)

	BeforeSuite(func() {
		cfg = config.Config{
			DBConnectTimeout: 30 * time.Second,
		}
	})

	Describe("Connecting to mongo DB", func() {
		It("Should connect and ping successfully", func() {
			_, err := db.Connect(cfg)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
