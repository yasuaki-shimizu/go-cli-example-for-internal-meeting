package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	price   int
	members []string
)

func main() {
	cmd := cobra.Command{
		Use:  "Calculate the members payment.",
		RunE: run,
	}
	cmd.Flags().IntVarP(&price, "price", "p", 0, "total price")
	cmd.Flags().StringSliceVarP(&members, "members", "m", nil, "members joined in party")

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	num := len(members)
	if num == 0 {
		return errors.New("no one joined in this party")
	}

	mod := price % num
	payment := price / num
	for i, member := range members {
		p := payment
		if i == 0 {
			p += mod
		}
		fmt.Printf("%s: %d\n", member, p)
	}

	return nil
}
