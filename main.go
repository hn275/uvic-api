package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hn275/uvic-api/cookies"
)

const (
	URL string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=CSC&txt_term=202301&txt_courseNumber=&startDatePicker=&endDatePicker=&pageOffset=0&pageMaxSize=1000&sortColumn=subjectDescription&sortDirection=asc"
)

func main() {
	sessionCookies := cookies.GetSessionCookie("202301")
	fmt.Println(sessionCookies)
	/*
		prints out:
			[
			JSESSIONID=4A8A7BDAA1C9FD9EFE9F9818F9780E40;
			Path=/StudentRegistrationSsb;
			HttpOnly;
			Secure;
			SameSite=Lax

			UVicPMember=!0vNLws/WbQTvacSlixZN9qU1tuJxXrQfdHWeGgPCgseyAtxMu9EQ+PQU0/ActuGVOqC3STEfHhN26Ek=;
			Path=/;
			HttpOnly;
			Secure
			]
	*/

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{}

	for _, cookie := range sessionCookies {
		req.AddCookie(cookie)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	buf, _ := io.ReadAll(res.Body)
	jsonRes := make(map[string]interface{})

	if err := json.Unmarshal(buf, &jsonRes); err != nil {
		panic(err)
	}

	fmt.Println(jsonRes)
}
