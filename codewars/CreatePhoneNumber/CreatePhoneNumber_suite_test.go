package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreatePhoneNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreatePhoneNumber Suite")
}
