// Package hwimpl includes implementation of microservice pkg/hw/HelloWorld
package hwimpl

import (
	"errors"
	"strings"

	"github.com/go-ozzo/ozzo-validation"

	"github.com/govinda-attal/hello-world/pkg/core/status"
	"github.com/govinda-attal/hello-world/pkg/hw"
)

// HelloWorld is an internal concrete implementation of pkg/hw/HelloWorld interface.
type HelloWorld struct {
	// Typically any dependencies
}

// New returns an instance of HelloWorld Implementation.
func New() *HelloWorld {
	return &HelloWorld{}
}

// Hello method returns a hello greetings message addressed to a person/bot with given name.
func (h *HelloWorld) Hello(name string) (*hw.HelloRs, error) {
	if err := h.validateRq(name); err != nil {
		return nil, err
	}
	rs := &hw.HelloRs{}
	rs.Greetings = strings.Join([]string{"Hello", name, "!"}, " ")
	return rs, nil
}

func (h *HelloWorld) validateRq(name string) error {
	notError := func(value interface{}) error {
		if s, _ := value.(string); strings.EqualFold(s, "error") {
			return errors.New("must not be error")
		}
		return nil
	}
	if err := validation.Validate(name, validation.By(notError)); err != nil {
		return status.ErrBadRequest.WithMessage(err.Error())
	}
	return nil
}
