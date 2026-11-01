package expenses

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
}

func OpenExpenses() []Expense {
	var exp []Expense

	allExp, err := os.ReadFile("./expenses.json")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil
		}
		os.Create("./expenses.json")

	}
	if len(allExp) > 0 {
		err = json.Unmarshal(allExp, &exp)
		if err != nil {
			return nil
		}
	}
	return exp
}
func write(expenses []Expense) {
	bytedExp, err := json.Marshal(expenses)
	if err != nil {
		fmt.Println("Fail to Marshal:", err)
		return
	}

	os.WriteFile("./expenses.json", bytedExp, 0644)
}
func AddExpense(description string, amount float64) {
	maxID := 0
	exp := OpenExpenses()
	for _, v := range exp {
		maxID = max(v.ID, maxID)
	}

	createdExpense := Expense{
		ID:          maxID + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}
	exp = append(exp, createdExpense)

	write(exp)

	fmt.Println("Expense added successfully, ID:", createdExpense.ID)

}

func UpdateExpense(id int, description string, amount float64) {
	allExpenses := OpenExpenses()
	found := false
	for i, v := range allExpenses {
		if v.ID == id {
			found = true
			allExpenses[i] = Expense{
				ID:          v.ID,
				Date:        v.Date,
				Description: description,
				Amount:      amount,
			}
		}
	}
	if !found {
		fmt.Println("Expense with id", id, "not exists")
		return
	}
	write(allExpenses)

	fmt.Println("Expense updated successfully, ID:", id)

}

func DeleteExpense(id int) {
	allExpenses := OpenExpenses()
	found := false

	var temp []Expense
	for i, v := range allExpenses {
		if v.ID != id {
			temp = append(temp, allExpenses[i])
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Expense with id", id, "not exists")
		return
	}
	write(temp)

	fmt.Println("Expense deleted successfully")
}

func ViewAllExpenses() {
	allExpenses := OpenExpenses()
	if len(allExpenses) == 0 {
		fmt.Println("Do not have expenses")
		return
	}

	for _, v := range allExpenses {
		fmt.Println("ID:", v.ID)
		fmt.Println("Date:", v.Date)
		fmt.Println("Description:", v.Description)
		fmt.Println("Amount:", v.Amount)
		fmt.Println("----------------------------------------------")
	}
}

func ViewSummary() {
	allExpenses := OpenExpenses()
	var summary float64
	for _, v := range allExpenses {
		summary += v.Amount
	}
	fmt.Println("Total expenses:", summary)
}
