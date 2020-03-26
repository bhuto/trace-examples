// Code generated by goa v3.1.1, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen github.com/DataDog/trace-examples/go/goa-basic/design -o
// $(GOPATH)/src/github.com/DataDog/trace-examples/go/goa-basic

package server

import (
	calc "github.com/DataDog/trace-examples/go/goa-basic/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v
}
