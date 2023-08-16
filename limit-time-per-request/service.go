package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

type Service struct {
	AllowedDuration time.Duration
	BaseUrl         string
	User            *User
}

func (s *Service) Process(path string) error {
	u, err := url.JoinPath(s.BaseUrl, path)
	if err != nil {
		return fmt.Errorf("can't join url path %w", err)
	}

	resp, err := http.Get(u)
	if err != nil {
		return fmt.Errorf("can't send request %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("wrong response %s", resp.Status)
	}

	defer func() { _ = resp.Body.Close() }()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return writeToFile("output.txt", data, path)
}

func NewService(sec time.Duration, baseUrl string) *Service {
	return &Service{AllowedDuration: sec, BaseUrl: baseUrl}
}

type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64
}

func writeToFile(fname string, data []byte, urlPath string) error {
	f, err := os.OpenFile(path.Join("xxx", fname),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0775)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, _ = fmt.Fprintf(f, "endpoint: %s\n", urlPath)

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
