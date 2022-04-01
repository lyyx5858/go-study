package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

const limit = 25
const total = 100
const page = total / limit

func fetch(url string, ch chan string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	client := &http.Client{}
	rsp, _ := client.Do(req)
	defer rsp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(rsp.Body)
	doc.Find(".item").Each(func(_ int, s *goquery.Selection) {
		num := s.Find(".pic em").Text()
		title := s.Find("span.title::first-child").Text()
		ch <- fmt.Sprintf("top %s: %s\n", num, title)
	})
}

func main() {
	ch := make(chan string)
	for i := 0; i < page; i++ {
		start := i * limit
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(start)
		go fetch(url, ch)
	}
	for i := 0; i < total; i++ {
		top := <-ch
		fmt.Println(top)
	}
}
