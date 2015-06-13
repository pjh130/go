package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/spider"
	//	"net/url"
	"strings"
)

type StarsInfo struct {
	Href        string
	Description string
}

var info []StarsInfo

type StarsProcesser struct {
}

func NewStarsProcesser() *StarsProcesser {
	return &StarsProcesser{}
}

func (this *StarsProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	query.Find(`li[class="repo-list-item public source"]`).Each(getStarsItem)

	//也许里边包含有分支
	query.Find(`li[class="repo-list-item public fork"]`).Each(getStarsItem)

	//	bFind = false
	query.Find(`div[class="pagination"]`).Find("a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if "Next" == text {
			next_page_href, _ := s.Attr("href")
			if next_page_href == "" {
				//				fmt.Println("next_page_href not find ")
				//				p.SetSkip(true)
			} else {
				fmt.Println("next_page_href: ", next_page_href)
				p.AddTargetRequest(next_page_href, "html")
				//		p.AddTargetRequestWithHeaderFile("http://weixin.sogou.com/weixin"+next_page_href, "html", "weixin.sogou.com.json")
			}
		}
	})
}

func getStarsItem(i int, s *goquery.Selection) {

	//找出域名
	//	rq := p.GetRequest().GetUrl()
	//	l, _ := url.ParseRequestURI(rq)
	//	host := l.Host
	//	index := strings.Index(rq, host)
	//	if index != -1 {
	//		href = rq[0:index] + host + href
	//	}

	//找出stars的链接
	href, _ := s.Find("h3").Find("a").Attr("href")
	href = "https://github.com" + href
	fmt.Printf("Href: %v \r\n", href)

	//找出描述
	description := strings.TrimSpace(s.Find(`p[class="repo-list-description"]`).Text())
	fmt.Printf("description: %v \r\n", description)

	if len(href) > 0 {
		var add StarsInfo
		add.Href = href
		add.Description = description
		info = append(info, add)
	}
}

type StarsPipelineConsole struct {
}

func (this *StarsPipelineConsole) Process(items *page_items.PageItems, t com_interfaces.Task) {

}

func GetStars(userName string) []StarsInfo {
	spider.NewSpider(NewStarsProcesser(), "GetStars").
		AddUrl("https://github.com/stars/"+userName, "html").
		AddPipeline(&StarsPipelineConsole{}). // Print result on screen
		SetThreadnum(3).                      // Crawl request by three Coroutines
		Run()

	fmt.Println("hrefs length: ", len(info))
	return info
}
