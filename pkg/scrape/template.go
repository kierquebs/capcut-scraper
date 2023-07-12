package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/kierquebs/capcut-scraper/util"
)


func GetRenderData(config util.Config,id int64) (string, error)  {

	var data string
	var err error

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("script[id=RENDER_DATA]", func(e *colly.HTMLElement) {
		data = e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := fmt.Sprintf("%s%d",config.TemplateURL,id)
	if err = c.Visit(url); err != nil {
		return data, err
	}
	return data, nil

}

func GetRenderDetail(config util.Config,id int64) (string, error)  {

	var data string
	var err error

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("script:contains('_ROUTER_DATA')", func(e *colly.HTMLElement) {
		text := e.Text
		data = util.StringReplace(text,"window._ROUTER_DATA = ", "", -1)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := fmt.Sprintf("%s%d",config.TemplateDetailURL,id)
	if err = c.Visit(url); err != nil {
		return data, err
	}
	return data, nil

}