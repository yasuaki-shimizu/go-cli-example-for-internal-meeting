package main

import (
	"cliexample/yakiniku/payment"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type executor struct {
	price   int
	members []string
	cmd     *cobra.Command
}

func newExecutor() *executor {
	e := executor{}

	cmd := cobra.Command{
		Use:  "Calculate the members payment.",
		RunE: e.run,
	}
	cmd.Flags().IntVarP(&e.price, "price", "p", 0, "total price")
	cmd.Flags().StringSliceVarP(&e.members, "members", "m", nil, "members joined in party")

	e.cmd = &cmd
	return &e
}

func (e *executor) Execute() error {
	return e.cmd.Execute()
}

func (e *executor) run(cmd *cobra.Command, args []string) error {
	calc, err := payment.NewCalculator(e.members, e.price)
	if err != nil {
		return err
	}

	responsibilities := calc.Calculate()
	for _, r := range responsibilities {
		fmt.Println(r)
	}

	return nil
}

func main() {
	exe := newExecutor()

	err := exe.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
