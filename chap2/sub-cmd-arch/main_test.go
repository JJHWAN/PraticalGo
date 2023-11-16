package main

import (
	"bytes"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSubCmdArch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SubCmdArch Suite")
}

var (
	usageMessage = `Usage: mync [http|grpc] -h

http: A HTTP Client

http: <options> server

Options: 
  -verb string
    	HTTP method (default "GET")


grpc A gRPC client.

grpc: <options> server

Options: 
  -body string
    	Body of request
  -method string
    	Method to call
`
)

type testConfig struct {
	args   []string
	output string
	err    error
}

var _ = Describe("SubCmdArch Test", func() {

	var byteBuf *bytes.Buffer

	BeforeEach(func() {
		byteBuf = new(bytes.Buffer)
	})

	AfterEach(func() {
		byteBuf.Reset()
	})

	Context("with empty argument", func() {
		config := testConfig{
			args:   []string{},
			err:    errInvalidSubCommand,
			output: "Invalid sub-command specified\n" + usageMessage,
		}
		It("should return Invalid sub-command error", func() {
			err := handleCommand(byteBuf, config.args)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(config.err))
			Expect(byteBuf.String()).To(Equal(config.output))
		})
	})

	Context("with -h flag", func() {
		config := testConfig{
			args:   []string{"-h"},
			err:    nil,
			output: usageMessage,
		}
		It("should return Invalid sub-command error", func() {
			err := handleCommand(byteBuf, config.args)
			Expect(err).ToNot(HaveOccurred())
			Expect(err).To(BeNil())
			Expect(byteBuf.String()).To(Equal(config.output))
		})
	})

	Context("with invalid sub command", func() {
		config := testConfig{
			args:   []string{"foo"},
			err:    errInvalidSubCommand,
			output: "Invalid sub-command specified\n" + usageMessage,
		}
		It("should return Invalid sub-command error", func() {
			err := handleCommand(byteBuf, config.args)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(config.err))
			Expect(byteBuf.String()).To(Equal(config.output))
		})
	})
})
