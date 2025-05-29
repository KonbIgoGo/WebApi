package handlers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/project/webapi/internal/entity"
)

type bigencHandler struct {
	url string
}

func (h *bigencHandler) GetArticles() (entity.Article, error) {
	data, err := readRaw(h.url)

	if err != nil {
		return entity.Article{}, err
	}

	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return entity.Article{}, err
	}

	resArticle := entity.Article{}

	doc.Find("h1.bre-article-header-title").First().Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		resArticle.Title = text
	})

	doc.Find("div.bre-article-body").Not("span.bre-media-image _note-exclude").Each(func(i int, s *goquery.Selection) {
		cloned := s.Clone()
		cloned.Find("span._note-exclude").Each(func(_ int, ex *goquery.Selection) {
			ex.Remove()
		})
		text := strings.TrimSpace(cloned.Text())
		resArticle.Data = text
	})

	return resArticle, nil
}
