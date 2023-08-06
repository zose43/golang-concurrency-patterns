package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type arg struct {
	phoneMap map[string]string
	u        User
}

func (a arg) phones() map[string]string {
	return a.phoneMap
}

func (a arg) user() User {
	return a.u
}

type jobArg interface {
	phones() map[string]string
	user() User
}

type User struct {
	Id    int    `id:"number"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string
}

func addPhone(ctx context.Context, arg interface{}) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("cancell %w", ctx.Err())
	default:
		time.Sleep(500 * time.Millisecond)
		jobArg, ok := arg.(jobArg)
		if !ok {
			return nil, fmt.Errorf("can't match argument interface")
		}

		user := jobArg.user()
		phones := jobArg.phones()
		phone, ok := phones[user.Name]
		if !ok {
			return nil, fmt.Errorf("can't match phone number by user name")
		}

		user.Phone = phone
		return user, nil
	}
}

func makeJobs(users []User, phones map[string]string) []Job {
	const jType = "clients"
	jobsBulk := make([]Job, len(users))
	for i, user := range users {
		arg := arg{u: user, phoneMap: phones}
		job := NewJob(&JobDescriptor{
			Type: jType,
		},
			addPhone,
			arg)
		jobsBulk[i] = job
	}
	return jobsBulk
}

func main() {
	users, phones, err := openDataset()
	if err != nil {
		log.Fatal(err)
	}

	wp := NewWp(5)
	jobsBulk := makeJobs(users, phones)
	go wp.GenerateFrom(jobsBulk)

	go wp.Run(context.Background())

	outputs := make([]Result, 0)
	for res := range wp.Results() {
		log.SetPrefix("error: ")
		if res.Err != nil {
			log.Println(res.Err)
			continue
		}
		outputs = append(outputs, res)
	}

	fmt.Printf("done, job-type = %s, all count = %d", outputs[0].Descriptor.Type, len(outputs))
}
