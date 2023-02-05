package uvicapi

import (
	"bytes"
	"net/url"
)

func (c *UVicAPI) GetTerms() ([]byte, error) {
	u, err := url.Parse(BASE + "/classSearch/getTerms")
	if err != nil {
		return nil, err
	}

	setQuery(u, map[string]string{
		"searchTerms": "",
		"offset":      "0",
		"max":         "500",
	})

	res, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(res.Body); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
