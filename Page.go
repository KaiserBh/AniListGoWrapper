package anilistgo

import (
	"encoding/json"
	"fmt"
)

// Page Page of data
type Page struct {
	PageInfo PageInfo `json:"pageInfo"`
	Media    []Media  `json:"media"`
}

// PageInfo represents the information regarding the Page of Anilist query
type PageInfo struct {
	Total       int  `json:"total"`
	PerPage     int  `json:"perPage"`
	CurrentPage int  `json:"currentPage"`
	LastPage    int  `json:"lastPage"`
	HasNextPage bool `json:"hasNextPage"`
}

// NewPageQuery Create new NewMediaTrend Object
func NewPageQuery() *Page {
	p := Page{}
	return &p
}

// PaginationByTitle  search Anilist Media by title returns arrayList of Media objects, and pageInfor takes title string, page (which page to look for), PerPage The amount of entries per page, max 50
func (p *Page) PaginationByTitle(title string, page int, perPage int) (bool, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ Page(page: %d, perPage: %d) {
			 pageInfo {
      			total,
      			perPage,
      			currentPage,
      			lastPage,
     		 hasNextPage,
   			 },
			
			media(search: "%v") {
				%s
			}
		}
	}
	`, page, perPage, title, MediaQuery),
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return false, err
	}

	request, err := PostRequest(jsonValue)
	if err != nil {
		return false, err
	}

	cleanData := CleanPageJSON(request)
	if err := json.Unmarshal(cleanData, &p); err != nil {
		return false, err
	}

	return true, nil
}
