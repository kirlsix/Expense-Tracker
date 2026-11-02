<h1>Expense Tracker CLI</h1>

A simple command-line expense tracker for managing your personal finances. Add, update, delete, and view expenses, and get summaries for all time or for a specific month — all from your terminal.

Features
Add an expense with a description and amount
Update an existing expense
Delete an expense by ID
View all expenses
View a summary of all expenses
View a summary of expenses for a specific month

go version go1.27.0

Installation
```
git clone https://github.com/kirlsix/Expense-Tracker
cd expense-tracker
go build -o expense-tracker
```
Optionally, install it to your $GOPATH/bin:
``` 
go install
```
<h1>Usage</h1>
Add an expense

```
$ expense-tracker add --description "Lunch" --amount 20
Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
Expense added successfully (ID: 2)
```
List all expenses
```bash
$ expense-tracker list

ID: 1
Date: 2026-11-02 ...
Description: Lunch
Amount: 20
----------------------------------------------
ID: 2
Date: 2026-11-02 ...
Description: Dinner
Amount: 10
----------------------------------------------
```
Update an expense
```
$ expense-tracker update --id 1 --description "Lunch with colleagues" --amount 25
Expense updated successfully
```
Delete an expense

```
$ expense-tracker delete --id 2
Expense deleted successfully
```
Summary of all expenses
```
$ expense-tracker summary
Total expenses: $30
```
Summary for a specific month
```
$ expense-tracker summary --month 8
Total expenses for August: $20
```
License<br>
MIT License — free to use and modify.

Author<br>
<a href="https://github.com/kirlsix">@kirlsix</a>
