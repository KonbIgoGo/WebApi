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
	if strings.Contains(url, "bigenc.ru") {
		return &bigencHandler{
			url: url,
		}, nil
	}

	return nil, errors.ErrUnsupported
}
