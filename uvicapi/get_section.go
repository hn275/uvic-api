package uvicapi

import (
	"bytes"
	"net/url"
	"strconv"
)

// NOTE:
// ignoring `queryParams.Term`, since UVic server won't sent back anything that's
// not in the term set by session (called in `NewAPI(term string)`)
func (c *UVicAPI) GetSection(queryParams UVicQueryParams) ([]byte, error) {
	getUrl, err := url.Parse(BASE + "searchResults/searchResults")
	if err != nil {
		return nil, err
	}

	q := map[string]string{
		"txt_subject":      queryParams.Subject,
		"txt_term":         c.Term,
		"txt_courseNumber": queryParams.CourseNumber,
		"startDatepicker":  "",
		"endDatepicker":    "",
		"pageOffset":       strconv.Itoa(queryParams.Offset),
		"pageMaxSize":      strconv.Itoa(queryParams.Max),
		"sortColumn":       "subjectDescription",
		"sortDirection":    "asc",
	}

	setQuery(getUrl, q)

	res, err := (*c).Get(getUrl.String())
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
