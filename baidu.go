package main

import (
"context"
"fmt"
"github.com/PuerkitoBio/goquery"
"github.com/chromedp/chromedp"
"log"
"strings"
"time"
)

func main() {
	fmt.Println("==========================================")
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("disable-gnu", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c0, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	c1, cancel := chromedp.NewContext( c0, chromedp.WithLogf(log.Printf))
	defer cancel()
	// chromedp.Run(c1, make([]chromedp.Action, 0, 1)...)
	// 给每个页面的爬取设置超时时间 10s后关闭
	c2, cancel := context.WithTimeout(c1, 20 * time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(c2,
		chromedp.Navigate("https://www.baidu.com"),
		chromedp.WaitVisible(`#wrapper`, chromedp.ByID),
		chromedp.OuterHTML(`document.querySelector("html")`, &htmlContent, chromedp.ByJSPath),
		//chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
		chromedp.SendKeys("#kw", "test\n", chromedp.ByID),
		chromedp.Sleep(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(htmlContent)

	//已经将网页存放在htmlContent，下面开始分析
	//doc =htmlContent
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("strong[class=J-p-209954]").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})


	//doc.Find("span").Each(func(i int, s *goquery.Selection) {
	//	// fmt.Println(enc.ConvertString(s.Text()))
	//	//	fmt.Println(s.Text())
	//	s1:=s.Text()
	//	// fmt.Println(strings.Index(s1, ".00"))
	//	if strings.Index(s1,".00")>0 {
	//		fmt.Println(enc.ConvertString(s1[165:181]))
	//	}
	//})

	fmt.Println("===========================================")

	//fmt.Println(htmlContent)
}