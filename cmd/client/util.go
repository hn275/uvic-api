package client

import "net/url"

func setQuery(u *url.URL, q map[string]string) {
	s := make(url.Values)
	for k, v := range q {
		s[k] = []string{v}
	}

	u.RawQuery = s.Encode()
}
