package cmd

import (
	expenses "ExpenceTracker/Expenses"

	"github.com/spf13/cobra"
)

// updateExpenseCmd represents the updateExpense command
var updateExpenseCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the information about an expense.",
	Long:  `A function that updates the information about an expense. It takes an ID, a description, and an amount as input.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		description, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetFloat64("amount")

		expenses.UpdateExpense(id, description, amount)
	},
}

func init() {
	rootCmd.AddCommand(updateExpenseCmd)
	updateExpenseCmd.Flags().Int("id", 0, "The ID of the expense being modified.")
	updateExpenseCmd.Flags().String("description", "", "The description of the expense.")
	updateExpenseCmd.Flags().Float64("amount", 0, "The amount of the expense.")
}
