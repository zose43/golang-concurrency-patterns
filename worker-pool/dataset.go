package main

import (
	"encoding/json"
	"fmt"
	"golang-concurrency-patterns/lib/storage"
	"log"
	"os"
	"path"
)

type dataset struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone_number"`
}

type client struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Email string `json:"email"`
}

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

func convert() {
	bytes, err := os.ReadFile(path.Join(storage.Dir, "dataset.json"))
	if err != nil {
		log.Fatal(err)
	}

	var elems []dataset
	if err = json.Unmarshal(bytes, &elems); err != nil {
		log.Fatal(err)
	}

	clients := make([]client, len(elems))
	phones := make(map[string]string, len(elems))
	for i, elem := range elems {
		phones[elem.Name] = elem.Phone
		clients[i] = client{Name: elem.Name, Id: elem.Id, Email: elem.Email}
	}

	f, err := os.Create(path.Join(storage.Dir, "clients.json"))
	if err != nil {
		log.Fatal(err)
	}
	clientsBt, _ := json.Marshal(clients)
	_, err = f.Write(clientsBt)
	if err != nil {
		log.Print("can't save clients")
	}

	f, err = os.Create(path.Join(storage.Dir, "phone_numbers.json"))
	if err != nil {
		log.Fatal(err)
	}
	phonesBt, _ := json.Marshal(phones)
	_, err = f.Write(phonesBt)
	if err != nil {
		log.Print("can't save phones")
	}
}
