package cmd

import (
	expenses "ExpenceTracker/Expenses"
	"time"

	"github.com/spf13/cobra"
)

// summaryExpensesCmd represents the summaryExpenses command
var summaryExpensesCmd = &cobra.Command{
	Use:   "summary",
	Short: "With the --month flag, outputs expenses only for the specified month; without the flag, outputs all expenses.",
	Long: `A function that outputs all expenses. If the --month flag is specified, 
	it outputs only the expenses for the specified month; 
	without the flag, it outputs all expenses.`,
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")
		if month == 0 {
			expenses.ViewSummary()
		} else {
			expenses.ViewSummaryForMonth(time.Month(month))
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryExpensesCmd)
	summaryExpensesCmd.Flags().Int("month", 0, "month")
}
