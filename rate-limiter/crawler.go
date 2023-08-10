package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Crawler(link string) (string, []string, error) {
	body, err := grabPage(link)
	if err != nil {
		return "", nil, err
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", nil, fmt.Errorf("can't open body %s %w", link, err)
	}
	_ = body.Close()

	title, links, err := scrapPage(doc, link)
	if err != nil {
		return "", nil, err
	}
	return title, links, nil
}

func scrapPage(doc *goquery.Document, link string) (title string, links []string, err error) {
	title = doc.Find("title").Text()
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		{
			val, ok := s.Attr("href")
			if ok {
				// todo only link by root domen, fx path
				val, err = checkUrl(val, link)
				links = append(links, val)
			}
		}
	})
	return
}

func checkUrl(val string, root string) (string, error) {
	if !strings.Contains(val, "http") {
		full, err := url.JoinPath(root, val)
		if err != nil {
			return "", fmt.Errorf("can't join url %w", err)
		}
		return full, nil
	}
	return val, nil
}

func grabPage(link string) (io.ReadCloser, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("can't get page %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("can't get page %s, %s", link, resp.Status)
	}
	return resp.Body, nil
}
