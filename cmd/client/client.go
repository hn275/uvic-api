package client

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	http.Client
}

const BASE = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/"
const MAX_SIZE = 500

func NewClient() (*Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	c := Client{}
	c.Client.Jar = cookieJar

	return &c, nil
}

func (c *Client) SetTerm(term string) error {
	requestUrl, err := url.Parse(BASE + "term/search")
	if err != nil {
		return err
	}

	setQuery(requestUrl, map[string]string{"mode": "search"})

	contentType := "application/x-www-form-urlencoded"
	reqBody := bytes.NewBufferString("term=" + term)

	_, err = (*c).Post(requestUrl.String(), contentType, reqBody)
	if err != nil {
		return err
	}

	return nil
}
