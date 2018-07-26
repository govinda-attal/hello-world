package mocks

import (
	"github.com/govinda-attal/hello-world/pkg/hw"
)

// HelloWorldMock is a mock helloworld microservice.
type HelloWorldMock struct {
	HelloCall struct {
		Receives struct {
			Name string
		}
		Returns struct {
			HelloRs *hw.HelloRs
			Error   error
		}
	}
}

// Hello Mock.
func (hm *HelloWorldMock) Hello(name string) (*hw.HelloRs, error) {
	hm.HelloCall.Receives.Name = name
	return hm.HelloCall.Returns.HelloRs, hm.HelloCall.Returns.Error
}
