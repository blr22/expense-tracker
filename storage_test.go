package main

import (
	"path/filepath"
	"testing"
)

func TestNewStorage(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "data")

	_, err := NewStorage(path)
	if err != nil {
		t.Fatalf("failed to create storage %q: %v", path, err)
	}
}
