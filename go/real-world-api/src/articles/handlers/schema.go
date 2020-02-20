package handlers

import (
	userHanlders "real-world-api/src/users/handlers"
	"time"
)

// ArticleSchema godoc
type ArticleSchema struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Description string `json:"description"`
	Body string `json:"body"`
	TagList []string `json:"TagList"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Favorited bool `json:"favorited"`
	FavoritesCount uint `json:"favoritesCount"`
	Author userHanlders.ProfileSchema
}