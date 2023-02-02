package uvicapi

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type UVicAPI struct {
	http.Client
	QueryParam UVicQueryParams
}

type UVicQueryParams struct {
	Subject      string // ie: "CSC"
	Term         string // ie: "202301"
	CourseNumber string // ie: "225"
	Offset       int
	Max          int // max 500
}

type UVicAPIError struct {
	error
	StatusCode int
}

const BASE = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/"
const MAX_SIZE = 0

func NewAPI(term string) (*UVicAPI, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	var c UVicAPI
	c.Jar = cookieJar

	requestUrl, err := url.Parse(BASE + "term/search")
	if err != nil {
		return nil, err
	}

	setQuery(requestUrl, map[string]string{"mode": "search"})

	_, err = c.PostForm(requestUrl.String(), url.Values{"term": []string{term}})
	if err != nil {
		return nil, err
	}

	c.QueryParam.Term = term // set term for future queries

	return &c, nil
}
