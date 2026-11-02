package cmd

import (
	expenses "ExpenceTracker/Expenses"

	"github.com/spf13/cobra"
)

// listAllExpensesCmd represents the listAllExpenses command
var listAllExpensesCmd = &cobra.Command{
	Use:   "list",
	Short: "Outputs all expenses.",
	Long:  `A function that outputs all expenses. It takes no input.`,
	Run: func(cmd *cobra.Command, args []string) {
		expenses.ViewAllExpenses()
	},
}

func init() {
	rootCmd.AddCommand(listAllExpensesCmd)
}
