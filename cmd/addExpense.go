package cmd

import (
	expenses "ExpenceTracker/Expenses"

	"github.com/spf13/cobra"
)

// addExpenseCmd represents the addExpense command
var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "A function that adds an expense to the general list.",
	Long:  `A function whose purpose is to add a newly created expense entry to the overall list of expenses.`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetFloat64("amount")
		description, _ := cmd.Flags().GetString("description")
		expenses.AddExpense(description, amount)

	},
}

func init() {
	rootCmd.AddCommand(addExpenseCmd)
	addExpenseCmd.Flags().Float64("amount", 0, "amount of the expense")
	addExpenseCmd.Flags().String("description", "", "description of the expense")
}
