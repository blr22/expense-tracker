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
