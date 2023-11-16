package cmd_test

import (
	"bytes"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/JJHWAN/chap2/sub-cmd-arch/cmd"
)

type testConfig struct {
	args   []string
	output string
	err    error
}

func TestSubCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestSubCmd Suite")
}

var httpUsageMessage = `
http: A HTTP Client

http: <options> server

Options: 
  -verb string
    	HTTP method (default "GET")

`

var _ = Describe("HttpCmd", func() {

	var byteBuf *bytes.Buffer

	BeforeEach(func() {
		byteBuf = new(bytes.Buffer)
	})

	AfterEach(func() {
		byteBuf.Reset()
	})

	Context("with empty arguments", func() {
		config := testConfig{
			args: []string{},
			err:  ErrorNoServerSpecified,
		}
		It("should print usage string", func() {
			err := HandleHttp(byteBuf, config.args)
			Expect(err).To(Equal(config.err))
			Expect(byteBuf.Len()).To(Equal(0))
		})
	})
})

var grpcUsageMessage = `
grpc A gRPC client.

grpc: <options> server

Options: 
  -body string
    	Body of request
  -method string
    	Method to call
`

var _ = Describe("GrpcCmd", func() {
	var byteBuf *bytes.Buffer

	BeforeEach(func() {
		byteBuf = new(bytes.Buffer)
	})

	AfterEach(func() {
		byteBuf.Reset()
	})

	Context("with empty arguments", func() {
		config := testConfig{
			args:   []string{},
			output: "",
			err:    ErrorNoServerSpecified,
		}
		It("should print usage string", func() {
			err := HandleGrpc(byteBuf, config.args)
			Expect(err).To(Equal(config.err))
			Expect(byteBuf.Len()).To(Equal(0))
		})
	})

	Context("with server argument", func() {
		config := testConfig{
			args:   []string{"server"},
			output: "Executing grpc command.\n",
			err:    nil,
		}
		It("should print usage string", func() {
			err := HandleGrpc(byteBuf, config.args)
			Expect(err).To(BeNil())
			Expect(byteBuf.String()).To(Equal(config.output))
		})
	})
})
