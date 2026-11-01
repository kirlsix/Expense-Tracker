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

	bytedExp, err := json.Marshal(exp)
	if err != nil {
		fmt.Println("Fail to Marshal:", err)
		return
	}

	os.WriteFile("./expenses.json", bytedExp, 0644)
	fmt.Println("Expense added successfully, ID:", createdExpense.ID)
}
