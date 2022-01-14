package tests

import (
	"log"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sailsforce/togo-read-micro/internal/config"
)

var _ = Describe("Config", func() {

	BeforeEach(func() {
		// load custom env variables
		os.Setenv("DATABASE_URL", "postgres://postgres:test@localhost:5555/testdb")
		os.Setenv("LOG_LEVEL", "blah")

		// create config, without db
		if err := config.NewServiceConfig(false); err != nil {
			log.Fatalf("error setting up test config: %v\n", err)
		}
	})

	AfterEach(func() {
		os.Clearenv()
	})

	Describe("Checking config setup", func() {
		Context("Logger", func() {
			It("use default log level if err", func() {
				Expect(config.Logger.GetLevel().String()).To(Equal("info"))
			})
		})
	})
})
