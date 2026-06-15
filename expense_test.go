package main

import (
	"testing"
)

func TestAddExpense(t *testing.T) {
	data := []struct {
		desc   string
		amount int
		expID  int
	}{
		{"Breakfast", 8, 1},
		{"Lunch", 20, 2},
		{"Dinner", 10, 3},
		{"IPhone 16", 830, 4},
	}

	el := NewExpenseList()
	for _, v := range data {
		if id := el.AddExpense(v.desc, v.amount); id != v.expID {
			t.Errorf("expected id=%d, got %d for %s", v.expID, id, v.desc)
		}
	}
	if len(el.items) != len(data) {
		t.Errorf("expected to be added %d values, %d added", len(data), len(el.items))
	}
}

func newFilledExpenseList() *ExpenseList {
	data := []struct {
		desc   string
		amount int
	}{
		{"Breakfast", 8},
		{"Lunch", 20},
		{"Dinner", 10},
		{"IPhone 16", 830},
	}

	el := NewExpenseList()
	for _, v := range data {
		el.AddExpense(v.desc, v.amount)
	}
	return el
}

func TestUpdateDesc(t *testing.T) {
	updates := map[int]string{1: "Chips", 2: "Cleaning", 3: "Rent", 4: "Travel"}
	
	el := newFilledExpenseList()
	for id, desc := range updates {
		if err := el.UpdateDesc(id, desc); err != nil {
			t.Errorf("unexpected error for id %d: %v", id, err)
		}
	}

	for _, v := range el.items {
		if v.Desc != updates[v.ID] {
			t.Errorf("expected desc %q for id %d, got %q", updates[v.ID], v.ID, v.Desc)
		}
	}
}
