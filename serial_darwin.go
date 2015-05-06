package serial

import (
	"fmt"
)

var errNotImplemented = fmt.Errorf("not implemented")

type port struct {
}

func New() Port {
	return &port{}
}

func (_ *port) Open(*Config) error {
	return errNotImplemented
}

func (_ *port) Read([]byte) (int, error) {
	return 0, errNotImplemented
}

func (_ *port) Write([]byte) (int, error) {
	return 0, errNotImplemented
}

func (_ *port) Close() error {
	return errNotImplemented
}
