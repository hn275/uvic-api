package main

import (
	"log"
	"os"
	"time"

	"github.com/hn275/uvic-api/uvicapi"
)

func main() {
	startTime := time.Now()

	term := "202301"
	log.Printf("fetching uvic data for term %s\n", term)
	uvicClient, err := uvicapi.NewAPI(term)
	if err != nil {
		log.Fatal(err)
	}

	// GetSection
	{
		q := uvicapi.UVicQueryParams{
			Subject:      "MATH",
			CourseNumber: "101",
			Max:          100,
			Offset:       0,
			Term:         "202209",
		}

		res, err := uvicClient.GetSection(q)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile("./json_examples/GetSection.json", res, 0666)
		log.Printf("DONE [GetSection] in %v\n", time.Since(startTime))
	}

	// GetAllSections
	{
		res, err := uvicClient.GetAllSections(1)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile("./json_examples/GetAllSections.json", res, 0666)
		log.Printf("DONE [GetAllSections] in %v\n", time.Since(startTime))
	}

	// GetCourseDesc
	{
		res, err := uvicClient.GetCourseDesc("20747") // wtf?
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(res)) // huh?
		log.Printf("DONE [GetCourseDesc] in %v\n", time.Since(startTime))
	}

	// GetTerms
	{
		res, err := uvicClient.GetTerms(1, 500)

		if err != nil {
			panic(err)
		}
		os.WriteFile("./json_examples/GetTerms.json", res, 0666)
		log.Printf("DONE [GetTerms] in %v\n", time.Since(startTime))
	}

	log.Fatalf("DONE in %v\n", time.Since(startTime))
}

/* Aomi's code below: to fetch literally everything
func main() {
	log := log.New(os.Stderr, "", 0)

	if len(os.Args) != 2 {
		log.Fatalf(""+
			"Usage: %v <TERM>\n"+
			"TERM is of the form YYYYMM, like 202205\n",
			os.Args[0])
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	c := &http.Client{Jar: jar}

	postURL(c, "term/search", map[string]string{
		"mode": "search",
	}, "application/x-www-form-urlencoded", "term="+os.Args[1])

	i := 0
	sections := make([]JSONValue, 0)
	totalSectionCount := "unknown"

	log.Printf("\rProgress: %v/%v", 0, totalSectionCount)
	for {
		b, err := getURL(c, "searchResults/searchResults", map[string]string{
			"txt_term":    os.Args[1],
			"pageOffset":  strconv.Itoa(i * maxSize),
			"pageMaxSize": "500",
		})
		if err != nil {
			log.Println(err)
			break
		}

		jsonResponse, err := fastjson.ParseBytes(b)
		if err != nil {
			log.Println(err)
			break
		}

		totalSectionCount := strconv.Itoa(jsonResponse.GetInt("sectionsFetchedCount"))

		data := jsonResponse.Get("data").GetArray()
		count := len(data)
		if count == 0 {
			break
		}

		for _, v := range data {
			sections = append(sections, JSONValue{Value: v})
		}

		i++
		log.Printf("\rProgress: %v/%v", len(sections), totalSectionCount)
	}

	b, err := json.MarshalIndent(sections, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))

	os.WriteFile("./data.json", b, 0666)
}

func getURL(c *http.Client, path string, query map[string]string) ([]byte, error) {
	reqURL, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, err
	}

	setQuery(reqURL, query)

	resp, err := c.Get(reqURL.String())
	if err != nil {
		return nil, err
	}

	b := readBody(resp)
	return b, nil
}

func postURL(c *http.Client, path string, query map[string]string, contentType string, body string) ([]byte, error) {
	reqURL, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, err
	}

	setQuery(reqURL, query)

	resp, err := c.Post(reqURL.String(), contentType, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	b := readBody(resp)
	return b, nil
}

func readBody(resp *http.Response) []byte {
	b := bytes.Buffer{}
	b.ReadFrom(resp.Body)
	return b.Bytes()
}

func setQuery(u *url.URL, query map[string]string) {
	out := make(url.Values)
	for k, v := range query {
		out[k] = []string{v}
	}
	u.RawQuery = out.Encode()
}

// because fastjson does't implement json.Marshaller
type JSONValue struct {
	Value *fastjson.Value
}

func (j *JSONValue) MarshalJSON() ([]byte, error) {
	return j.Value.MarshalTo(nil), nil
}}

func getURL(c *http.Client, path string, query map[string]string) ([]byte, error) {
	reqURL, err := url.Parse(BASE + path)
	if err != nil {
		return nil, err
	}

	setQuery(reqURL, query)

	resp, err := c.Get(reqURL.String())
	if err != nil {
		return nil, err
	}

	b := readBody(resp)
	return b, nil
}

func startSession(c *http.Client) ([]byte, error) {
	reqURL, err := url.Parse(BASE + "term/search")
	if err != nil {
		return nil, err
	}

	setQuery(reqURL, map[string]string{"mode": "search"})

	contentType := "application/x-www-form-urlencoded"
	resp, err := c.Post(reqURL.String(), contentType, bytes.NewBufferString("term="+TERM))
	if err != nil {
		return nil, err
	}

	b := readBody(resp)
	return b, nil
}

func readBody(resp *http.Response) []byte {
	b := bytes.Buffer{}
	b.ReadFrom(resp.Body)
	return b.Bytes()
}

func setQuery(u *url.URL, query map[string]string) {
	out := make(url.Values)
	for k, v := range query {
		out[k] = []string{v}
	}
	u.RawQuery = out.Encode()
}

// because fastjson does't implement json.Marshaller
type JSONValue struct {
	Value *fastjson.Value
}

func (j *JSONValue) MarshalJSON() ([]byte, error) {
	return j.Value.MarshalTo(nil), nil
}
*/
