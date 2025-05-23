package handlers

import (
	"errors"
	"strings"

	"github.com/project/webapi/internal/entity"
)

type WebHandler interface {
	GetArticles() (entity.Article, error)
}

func New(url string) (WebHandler, error) {
	println("T")
	if strings.Contains(url, "bigenc.ru") {
		return &bigencHandler{
			url: url,
		}, nil
	} else if strings.Contains(url, "wikipedia.org") {
		return &wikipediaHandler{
			url: url,
		}, nil
	}

	return nil, errors.ErrUnsupported
}
