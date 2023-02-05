package uvicapi

import (
	"bytes"
	"net/url"
	"strconv"
)

// Get all sections of all classes, 500 entries per page, offset is required
func (c *UVicAPI) GetAllSections(offset int) ([]byte, error) {
	reqUrl, err := url.Parse(banner + "searchResults/searchResults")
	if err != nil {
		return nil, err
	}

	setQuery(reqUrl, map[string]string{
		"txt_term":    c.Term,
		"pageOffset":  strconv.Itoa(offset * 500),
		"pageMaxSize": strconv.Itoa(500),
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
