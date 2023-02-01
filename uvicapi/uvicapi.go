package uvicapi

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

type UVicAPI struct {
	http.Client
}

const BASE = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/"
const MAX_SIZE = 500

func NewClient() (*UVicAPI, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	c := UVicAPI{}
	c.Client.Jar = cookieJar

	return &c, nil
}

func (c *UVicAPI) GetSections(term string, offset int) ([]byte, error) {
	reqUrl, err := url.Parse(BASE + "searchResults/searchResults")
	if err != nil {
		return nil, err
	}

	setQuery(reqUrl, map[string]string{
		"txt_term":    term,
		"pageOffset":  strconv.Itoa((offset + 1) * MAX_SIZE),
		"pageMaxSize": strconv.Itoa(MAX_SIZE),
	})

	res, err := c.Get(reqUrl.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var buf bytes.Buffer

	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *UVicAPI) SetTerm(term string) error {
	requestUrl, err := url.Parse(BASE + "term/search")
	if err != nil {
		return err
	}

	setQuery(requestUrl, map[string]string{"mode": "search"})

	contentType := "application/x-www-form-urlencoded"
	reqBody := bytes.NewBufferString("term=" + term)

	_, err = c.Post(requestUrl.String(), contentType, reqBody)
	if err != nil {
		return err
	}

	return nil
}
