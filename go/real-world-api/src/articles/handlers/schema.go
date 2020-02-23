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
	TagList []string `json:"tagList"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	Favorited bool `json:"favorited"`
	FavoritesCount uint `json:"favoritesCount"`
	Author userHanlders.ProfileSchema `json:"author"`
}

// CommentSchema godoc
type CommentSchema struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Body string `json:"body"`
	Author userHanlders.ProfileSchema `json:"author"`
}