package cmd

import (
	expenses "ExpenceTracker/Expenses"

	"github.com/spf13/cobra"
)

// deleteExpenseCmd represents the deleteExpense command
var deleteExpenseCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an expense.",
	Long:  `A function that deletes an expense. It takes the ID of the expense as input.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		expenses.DeleteExpense(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteExpenseCmd)
	deleteExpenseCmd.Flags().Int("id", 0, "The ID of the expense being deleted.")
}
