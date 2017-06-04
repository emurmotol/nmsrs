package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"reflect"

	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/model"
)

func SearchIndex(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	searchResults := []model.SearchResult{}

	if q != "" {
		users := model.SearchUsers(q)
		// registrants := model.SearchRegistrants(q)
		// searchResults = append(users, registrants...)
		searchResults = users // temporary
	}
	rd.JSON(w, http.StatusOK, searchResults)
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	typ := r.URL.Query().Get("type")

	user := model.User{}
	// registrant := model.Registrant{}

	searchResults := []model.SearchResult{}

	if q != "" {
		switch typ {
		case reflect.TypeOf(user).Name():
			searchResults = model.SearchUsers(q)
			break
		case "Registrant": // reflect.TypeOf(registrant).Name()
			break
		case "":
			users := model.SearchUsers(q)
			// registrants := model.SearchRegistrants(q)
			// searchResults = append(users, registrants...)
			searchResults = users // temporary
			break
		default:
			// todo: flash error here invalid type
		}
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}

	pagination := &helper.Paginator{
		Page:     page,
		Limit:    limit,
		Count:    len(searchResults),
		Interval: interval,
		QueryURL: r.URL.Query(),
	}

	if page > pagination.PageCount() {
		pagination.Page = 1
	}
	data := make(map[string]interface{})
	data["title"] = "Search"
	data["authUser"] = authUser(r)
	data["q"] = q
	data["type"] = typ
	data["searchResults"] = searchResults[pagination.Offset():pagination.EndIndex()]
	data["pagination"] = helper.Pager{
		Markup:     template.HTML(pagination.String()),
		Count:      pagination.Count,
		StartIndex: pagination.StartIndex(),
		EndIndex:   pagination.EndIndex(),
	}
	rd.HTML(w, http.StatusOK, "search/index", data)
}
