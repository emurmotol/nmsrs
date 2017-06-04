package model

import (
	"reflect"
	"time"
)

type SearchResult struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	HasPhoto  bool      `json:"has_photo"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
}

func SearchUsers(q string) []SearchResult {
	searchResults := []SearchResult{}
	user := User{}

	for _, user := range user.Search(q) {
		searchResults = append(searchResults, SearchResult{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			HasPhoto:  user.HasPhoto,
			CreatedAt: user.CreatedAt,
			Type:      reflect.TypeOf(user).Name(),
		})
	}
	return searchResults
}
