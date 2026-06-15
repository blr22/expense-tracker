package main

import (
	"fmt"
	"time"
)

type Expense struct {
	ID     int
	Desc   string
	Amount int
	Date   time.Time
}

func NewExpense(id int, desc string, amount int) Expense {
	return Expense{
		ID:     id,
		Desc:   desc,
		Amount: amount,
		Date:   time.Now(),
	}
}

type ExpenseList struct {
	items []Expense
}

func NewExpenseList() *ExpenseList {
	return &ExpenseList{}
}

func (el *ExpenseList) AddExpense(desc string, amount int) int {
	id := el.nextID()
	expense := NewExpense(id, desc, amount)
	el.items = append(el.items, expense)
	return id
}

func (el *ExpenseList) UpdateDesc(id int, desc string) error {
	for i, v := range el.items {
		if v.ID == id {
			el.items[i].Desc = desc
			el.items[i].Date = time.Now()
			return nil
		}
	}
	return fmt.Errorf("incorrect ID: %d", id)
}

func (el *ExpenseList) UpdateAmount(id int, amount int) error {
	for i, v := range el.items {
		if v.ID == id {
			el.items[i].Amount = amount
			el.items[i].Date = time.Now()
			return nil
		}
	}
	return fmt.Errorf("incorrect ID: %d", id)
}

func (el *ExpenseList) DeleteExpense(id int) error {
	for i, v := range el.items {
		if v.ID == id {
			el.items = append(el.items[:i], el.items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("incorrect ID: %d", id)
}

func (el *ExpenseList) ListAll() []Expense {
	return append([]Expense{}, el.items...)
}

func (el *ExpenseList) Summary() int {
	var sum int
	for _, v := range el.items {
		sum += v.Amount
	}
	return sum
}

func (el *ExpenseList) SummaryByMonth(month int) (int, error) {
	if month < 1 || month > 12 {
		return 0, fmt.Errorf("incorrect month, expected: 1-12, got: %d", month)
	}
	m := time.Month(month)
	var sum int
	for _, v := range el.items {
		if v.Date.Month() == m {
			sum += v.Amount
		}
	}
	return sum, nil
}

func (el *ExpenseList) nextID() int {
	var id int
	for _, v := range el.items {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
