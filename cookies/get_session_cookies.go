package cookies

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	TERM_URL string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/term/search?mode=search"
)

func GetSessionCookie(t string) []*http.Cookie {
	var setTermBody bytes.Buffer
	term := map[string]string{"term": t}
	if err := json.NewEncoder(&setTermBody).Encode(&term); err != nil {
		panic(err)
	}

	res, err := http.Post(TERM_URL, "application/x-www-form-urlencoded", &setTermBody)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	return res.Cookies()
}
