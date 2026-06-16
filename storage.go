package main

import "os"

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
