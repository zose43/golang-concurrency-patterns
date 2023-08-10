package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var rootURL string

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

	title, links, err := scrapPage(doc, link, root(link))
	if err != nil {
		return "", nil, err
	}
	return title, links, nil
}

func scrapPage(doc *goquery.Document, link string, root string) (title string, links []string, err error) {
	title = doc.Find("title").Text()
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		{
			val, ok := s.Attr("href")
			if ok {
				val, err = checkUrl(val, link)
				if strings.Contains(val, root) {
					links = append(links, val)
				}
			}
		}
	})
	return
}

func checkUrl(link string, root string) (string, error) {
	if !strings.Contains(link, "http") {
		//todo fix path
		full, err := url.JoinPath(root, link)
		if err != nil {
			return "", fmt.Errorf("can't join url %w", err)
		}
		return unescapedUrl(full)
	}
	return unescapedUrl(link)
}

func unescapedUrl(link string) (string, error) {
	link, err := url.PathUnescape(link)
	if err != nil {
		return "", fmt.Errorf("can't path unescape %w", err)
	}
	return link, nil
}

func grabPage(link string) (io.ReadCloser, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, fmt.Errorf("can't get page %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("can't get page %s", resp.Status)
	}
	return resp.Body, nil
}

func root(link string) string {
	if len(rootURL) == 0 {
		rootURL = link
	}
	return rootURL
}
