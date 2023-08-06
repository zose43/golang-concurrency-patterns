package main

import (
	"encoding/json"
	"fmt"
	"golang-concurrency-patterns/lib/storage"
)

func openDataset() (users []User, phones map[string]string, err error) {
	userBt, err := storage.OpenData("clients.json")
	if err != nil {
		return nil, nil, fmt.Errorf("can't read file %w", err)
	}
	phoneBt, err := storage.OpenData("phone_numbers.json")
	if err != nil {
		return nil, nil, fmt.Errorf("can't read file %w", err)
	}

	err = json.Unmarshal(userBt, &users)
	if err != nil {
		return nil, nil, fmt.Errorf("can't unmarshall clients %w", err)
	}
	err = json.Unmarshal(phoneBt, &phones)
	if err != nil {
		return nil, nil, fmt.Errorf("can't unmarshall phones %w", err)
	}
	return
}
