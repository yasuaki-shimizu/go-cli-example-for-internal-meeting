package payment_test

import (
	"cliexample/yakiniku/payment"
	"testing"

	"github.com/matryer/is"
)

func TestResponsibility(t *testing.T) {
	is := is.New(t)

	r := payment.Responsibility{
		Name:  "foo",
		Value: 1,
	}

	actual := r.String()

	expected := "foo: 1"

	is.Equal(expected, actual)
}
