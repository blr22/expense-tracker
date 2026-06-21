package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

func printExpenses(e []Expense) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(tw, "ID\tDate\tDescription\tAmount")
	for _, v := range e {
		date := v.Date.Format("2006-01-02 15:04:05")
		fmt.Fprintf(tw, "%d\t%s\t%s\t%d\n", v.ID, date, v.Desc, v.Amount)
	}
	tw.Flush()
}

func main() {
	s, err := NewStorage("data")
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}
	defer s.Close()

	expenses, err := s.ReadAll()
	if err != nil {
		log.Fatalf("failed to read storage: %v", err)
	}
	expenseList := NewExpenseList()
	expenseList.Load(expenses)

	if len(os.Args[1:]) == 0 {
		log.Fatalf("no command provided")
	}
	args := os.Args[1:]
	switch args[0] {
	case "add":
		var desc string
		var amount int
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)

		addCmd.StringVar(&desc, "description", "expense", "expense description")
		addCmd.IntVar(&amount, "amount", 0, "expense amount")
		addCmd.Parse(args[1:])

		id := expenseList.AddExpense(desc, amount)
		err := s.WriteAll(expenseList.All())
		if err != nil {
			log.Fatalf("failed to write expenses: %v", err)
		}
		fmt.Printf("Add successfully (ID: %d)\n", id)
	case "update":
		var id int
		var amount int
		var desc string
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

		updateCmd.IntVar(&id, "id", 0, "expense id")
		updateCmd.IntVar(&amount, "amount", 0, "expense amount")
		updateCmd.StringVar(&desc, "description", "description", "expense description")
		updateCmd.Parse(args[1:])

		var descSet, amountSet bool
		updateCmd.Visit(func(f *flag.Flag) {
			if f.Name == "description" {
				descSet = true
			}
			if f.Name == "amount" {
				amountSet = true
			}
		})

		if descSet {
			err := expenseList.UpdateDesc(id, desc)
			if err != nil {
				log.Fatalf("failed to update description: %v", err)
			}
		}
		if amountSet {
			err := expenseList.UpdateAmount(id, amount)
			if err != nil {
				log.Fatalf("failed to update amount: %v", err)
			}
		}

		err := s.WriteAll(expenseList.All())
		if err != nil {
			log.Fatalf("failed to write expenses: %v", err)
		}
		fmt.Println("Expense updated successfully")
	case "delete":
		var id int
		delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		delCmd.IntVar(&id, "id", 0, "expense id")
		delCmd.Parse(args[1:])

		err := expenseList.DeleteExpense(id)
		if err != nil {
			log.Fatalf("failed to delete expense: %v", err)
		}

		err = s.WriteAll(expenseList.All())
		if err != nil {
			log.Fatalf("failed to write expenses: %v", err)
		}
		fmt.Println("Expense deleted successfully")
	case "summary":
		var month int
		sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		sumCmd.IntVar(&month, "month", 1, "expense month")
		sumCmd.Parse(args[1:])

		var monthSet bool
		sumCmd.Visit(func(f *flag.Flag) {
			if f.Name == "month" {
				monthSet = true
			}
		})

		if monthSet {
			sum, err := expenseList.SummaryByMonth(month)
			if err != nil {
				log.Fatalf("failed to get summary: %v", err)
			}
			fmt.Printf("Total expenses for %s: %d\n", time.Month(month).String(), sum)
		} else {
			fmt.Printf("Total expenses: %d\n", expenseList.Summary())
		}
	case "list":
		printExpenses(expenseList.All())
	default:
		log.Fatalf("unknown command: %s", args[0])
	}
}
