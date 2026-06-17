package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

type Storage struct {
	file *os.File
}

func NewStorage(filename string) (*Storage, error) {
	f, err := os.OpenFile(filename+".csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &Storage{file: f}, nil
}

func (s *Storage) Close() error {
	return s.file.Close()
}

func (s *Storage) ReadAll() ([]Expense, error) {
	_, err := s.file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(s.file)
	r.FieldsPerRecord = 4
	var res []Expense
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		desc := record[1]
		amount, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		date, err := time.Parse(time.RFC3339, record[3])
		if err != nil {
			return nil, err
		}

		res = append(res, Expense{
			ID:     id,
			Desc:   desc,
			Amount: amount,
			Date:   date,
		})
	}
	return res, nil
}

func (s *Storage) WriteAll(data []Expense) error {
	records := make([][]string, 0, len(data))

	for _, v := range data {
		id := strconv.Itoa(v.ID)
		desc := v.Desc
		amount := strconv.Itoa(v.Amount)
		date := v.Date.Format(time.RFC3339)
		records = append(records, []string{id, desc, amount, date})
	}

	err := s.file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = s.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	w := csv.NewWriter(s.file)
	err = w.WriteAll(records)
	if err != nil {
		return err
	}
	return nil
}
