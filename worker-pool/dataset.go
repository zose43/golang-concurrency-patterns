package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func openDataset() (users []User, phones map[string]string, err error) {
	userBt, err := os.ReadFile("clients.json")
	if err != nil {
		return nil, nil, fmt.Errorf("can't read file %w", err)
	}
	phoneBt, err := os.ReadFile("phone_numbers.json")
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
