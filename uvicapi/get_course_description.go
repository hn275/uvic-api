package uvicapi

import (
	"bytes"
	"net/url"
)

func (c *UVicAPI) GetCourseDesc(crn string) ([]byte, error) {
	u, err := url.Parse(BASE + "searchResults/getCourseDescription")
	if err != nil {
		return nil, err
	}

	body := url.Values{
		"term":                  []string{c.QueryParam.Term},
		"courseReferenceNumber": []string{crn},
	}

	res, err := (*c).PostForm(u.String(), body)
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
