package model

import (
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SearchResult struct {
	Id        bson.ObjectId `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	HasPhoto  bool          `json:"has_photo"`
	CreatedAt time.Time     `json:"created_at"`
	Type      string        `json:"type"`
}

func SearchUsers(q string) []SearchResult {
	searchResults := []SearchResult{}
	user := User{}

	for _, user := range user.Search(q) {
		searchResults = append(searchResults, SearchResult{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			HasPhoto:  user.HasPhoto,
			CreatedAt: user.CreatedAt,
			Type:      reflect.TypeOf(user).Name(),
		})
	}
	return searchResults
}
