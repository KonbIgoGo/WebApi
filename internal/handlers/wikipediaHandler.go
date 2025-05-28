package handlers

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/project/webapi/internal/entity"
)

type wikipediaHandler struct {
	url string
}

func (h *wikipediaHandler) GetArticles() (entity.Article, error) {
	data, _ := readRaw(h.url)

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		log.Fatal(err)
	}

	resArticle := entity.Article{}

	doc.Find("h1.firstHeading").First().Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		resArticle.Title = text
	})

	var sb strings.Builder
	doc.Find("div.mw-content-ltr p").Each(func(i int, s *goquery.Selection) {
		paragraph := strings.TrimSpace(s.Text())
		if paragraph != "" {
			sb.WriteString(paragraph)
			sb.WriteString("\n\n")
		}
	})
	resArticle.Data = strings.TrimSpace(sb.String())

	return resArticle, nil
}
