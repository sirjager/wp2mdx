package service

import (
	"github.com/rs/zerolog"
)

type WordressSite struct {
	siteUrl string
	perPage int

	HasData bool

	Collections   []string
	AllPosts      []Post
	AllPages      []Page
	AllCategories []Category
	AllTags       []Tag
	AllUsers      []User
	AllMedia      []Media

	logger zerolog.Logger
}

func NewWordpressSite(url string, logger zerolog.Logger) *WordressSite {
	return &WordressSite{
		siteUrl:     url,
		perPage:     50,
		logger:      logger,
		HasData:     false,
		Collections: []string{"posts", "media", "users", "pages", "tags", "categories"},
	}
}
