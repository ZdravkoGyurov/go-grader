package storage_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ZdravkoGyurov/go-grader/pkg/app/config"
	"github.com/ZdravkoGyurov/go-grader/pkg/storage"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Storage Suite")
}

var _ = Describe("Storage", func() {
	var (
		ctx context.Context
		cfg config.Config
	)

	BeforeSuite(func() {
		ctx = context.Background()
		cfg = config.Config{
			DatabaseURI:         "mongodb://localhost:27017",
			DatabaseName:        "grader",
			DBConnectTimeout:    10 * time.Second,
			DBDisconnectTimeout: 10 * time.Second,
		}
	})

	Describe("Connecting to mongo DB", func() {
		It("Should connect and ping successfully", func() {
			_, err := storage.New(ctx, cfg)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
