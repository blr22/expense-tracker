package main

import (
	"encoding/csv"
	"path/filepath"
	"testing"
	"time"
)

func TestNewStorage(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "data")

	s, err := NewStorage(path)
	if err != nil {
		t.Fatalf("failed to create storage %q: %v", path, err)
	}
	if err = s.Close(); err != nil {
		t.Fatalf("failed to close storage: %v", err)
	}
}

func TestReadAll(t *testing.T) {
	data := [][]string{
		{"1", "Groceries", "35", "2026-10-01T00:00:00Z"},
		{"2", "Bus ticket", "2", "2026-10-02T00:00:00Z"},
		{"3", "Coffee", "7", "2026-10-04T00:00:00Z"},
	}
	want := []Expense{
		{1, "Groceries", 35, time.Date(2026,
			time.October, 1, 0, 0, 0, 0, time.UTC)},
		{2, "Bus ticket", 2, time.Date(2026,
			time.October, 2, 0, 0, 0, 0, time.UTC)},
		{3, "Coffee", 7, time.Date(2026,
			time.October, 4, 0, 0, 0, 0, time.UTC)},
	}
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "data")

	s, err := NewStorage(path)
	if err != nil {
		t.Fatalf("failed to create storage %q: %v", path, err)
	}
	defer s.Close()

	r := csv.NewWriter(s.file)
	err = r.WriteAll(data)
	if err != nil {
		t.Fatalf("csv.Writer.WriteAll failed: %v", err)
	}

	records, err := s.ReadAll()
	if err != nil {
		t.Fatalf("failed to ReadAll: %v", err)
	}

	if len(want) != len(records) {
		t.Fatalf("want %d records, got %d", len(want), len(records))
	}
	for i, r := range records {
		if want[i].ID != r.ID || want[i].Desc != r.Desc ||
			want[i].Amount != r.Amount || !want[i].Date.Equal(r.Date) {
			t.Errorf("record %d mismatched: got: %+v, want %+v", r.ID, r, want[i])
		}
	}
}
