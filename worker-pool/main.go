package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type extendedUser struct {
	user User
	err  error
}

type User struct {
	Id    string `id:"number"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string
}

func addPhone(user User, phones map[string]string) (User, error) {
	time.Sleep(500 * time.Millisecond)
	phone, ok := phones[user.Name]
	if !ok {
		// todo handle err
	}
	user.Phone = phone
	return user, nil
}

func handleUser(wg *sync.WaitGroup, input <-chan User, output chan<- extendedUser, phones map[string]string) {
	defer wg.Done()
	for user := range input {
		user, err := addPhone(user, phones)
		extUser := extendedUser{user: user, err: err}
		output <- extUser
	}
}

func main() {
	users, phones, err := openDataset()
	if err != nil {
		log.Fatal(err)
	}

	inputCh := make(chan User)
	go func() {
		for i := range users {
			inputCh <- users[i]
		}
		close(inputCh)
	}()

	outputCh := make(chan extendedUser)
	const workerNum = 4
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < workerNum; i++ {
			wg.Add(1)
			go handleUser(&wg, inputCh, outputCh, phones)
		}
		wg.Wait()
		close(outputCh)
	}()

	outputs := make([]extendedUser, 0)
	for extUser := range outputCh {
		log.SetPrefix("error: ")
		if extUser.err != nil {
			log.Println(extUser.err)
			continue
		}
		outputs = append(outputs, extUser)
	}

	fmt.Printf("done, clients count = %d", len(outputs))
}
