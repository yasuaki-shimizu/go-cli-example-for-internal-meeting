package payment_test

import (
	"cliexample/yakiniku/payment"
	"testing"

	"github.com/matryer/is"
)

func TestCalculator(t *testing.T) {
	// Calculatorの基本機能のテスト

	is := is.New(t)

	calc, err := payment.NewCalculator(
		[]string{"taro", "jiro", "ooshioheihachiro"},
		3000,
	)

	is.NoErr(err)

	responsibilies := calc.Calculate()

	expected := []payment.Responsibility{
		{Name: "taro", Value: 1000},
		{Name: "jiro", Value: 1000},
		{Name: "ooshioheihachiro", Value: 1000},
	}

	is.Equal(expected, responsibilies)
}

func TestCalculatorMod(t *testing.T) {
	// 金額が割り切れない場合１番目の人があまりを払うテスト

	is := is.New(t)

	calc, err := payment.NewCalculator(
		[]string{"more", "less1", "less2"},
		3100,
	)

	is.NoErr(err)

	responsibilies := calc.Calculate()

	expected := []payment.Responsibility{
		{Name: "more", Value: 1034}, // more pays more
		{Name: "less1", Value: 1033},
		{Name: "less2", Value: 1033},
	}

	is.Equal(expected, responsibilies)
}
