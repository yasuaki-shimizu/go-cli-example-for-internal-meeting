package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

type executer struct {
	price   int
	members []string
	cmd     *cobra.Command
}

func newExecuter() *executer {
	e := executer{}

	cmd := cobra.Command{
		Use:  "Calculate the members payment.",
		RunE: e.run,
	}
	cmd.Flags().IntVarP(&e.price, "price", "p", 0, "total price")
	cmd.Flags().StringSliceVarP(&e.members, "members", "m", nil, "members joined in party")

	e.cmd = &cmd
	return &e
}

func (e *executer) Execute() error {
	return e.cmd.Execute()
}

func (e *executer) run(cmd *cobra.Command, args []string) error {
	num := len(e.members)
	if num == 0 {
		return errors.New("no one joined in this party")
	}

	mod := e.price % num
	payment := e.price / num
	for i, member := range e.members {
		p := payment
		if i == 0 {
			p += mod
		}
		fmt.Printf("%s: %d\n", member, p)
	}

	return nil
}

func main() {
	exe := newExecuter()

	err := exe.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
