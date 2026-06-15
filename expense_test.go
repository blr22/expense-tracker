package main

import (
	"testing"
	"time"
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

func TestUpdateAmount(t *testing.T) {
	updates := map[int]int{1: 25, 2: 30, 3: 40, 4: 790}

	el := newFilledExpenseList()
	for id, amount := range updates {
		if err := el.UpdateAmount(id, amount); err != nil {
			t.Errorf("unexpected error for id %d: %v", id, err)
		}
	}

	for _, v := range el.items {
		if v.Amount != updates[v.ID] {
			t.Errorf("expected amount %d for id %d, got %d", updates[v.ID], v.ID, v.Amount)
		}
	}
}

func TestDeleteExpense(t *testing.T) {
	deletedIDs := map[int]bool{2: true, 1: true, 3: true}

	el := newFilledExpenseList()
	for id := range deletedIDs {
		if err := el.DeleteExpense(id); err != nil {
			t.Errorf("unexpected error for id %d", id)
		}
	}

	for _, v := range el.items {
		if deletedIDs[v.ID] {
			t.Errorf("id %d should be deleted", v.ID)
		}
	}
}

func TestListAll(t *testing.T) {
	el := newFilledExpenseList()
	l := el.ListAll()
	if len(el.items) != len(l) {
		t.Errorf("expected %d items, got %d", len(el.items), len(l))
	}
	el.items[0] = Expense{}
	if l[0].Desc == "" {
		t.Error("ListAll returned the same underlying array, expected copy")
	}
}

func TestSummary(t *testing.T) {
	el := newFilledExpenseList()
	var expSum int
	for _, v := range el.items {
		expSum += v.Amount
	}

	if sum := el.Summary(); expSum != sum {
		t.Errorf("expected sum %d, got %d", expSum, sum)
	}
}

func summaryByMonth(el *ExpenseList, month int) int {
	var sum int
	m := time.Month(month)
	for _, v := range el.items {
		if v.Date.Month() == m {
			sum += v.Amount
		}
	}
	return sum
}

func TestSummaryByMonth(t *testing.T) {
	invalidMonths := []int{-1, 22, 0, 13, 99}

	el := newFilledExpenseList()
	for _, month := range invalidMonths {
		if _, err := el.SummaryByMonth(month); err == nil {
			t.Errorf("month must be 1-12, not %d", month)
		}
	}

	for month := 1; month <= 12; month++ {
		expSum := summaryByMonth(el, month)
		if sum, _ := el.SummaryByMonth(month); sum != expSum {
			t.Errorf("expected summary by %s equal %d, got %d", time.Month(month), expSum, sum)
		}
	}
}
