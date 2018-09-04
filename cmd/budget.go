package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(budgetCmd)
}

var budgetCmd = &cobra.Command{
	Use:     "budget",
	Aliases: []string{"budg", "bud"},
	Short:   "Manage budget",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}
