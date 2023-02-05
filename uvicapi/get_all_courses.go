package uvicapi

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/valyala/fastjson"
)

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

const kuali string = "https://uvic.kuali.co/api/v1/catalog/courses/5d9ccc4eab7506001ae4c225"

func (c *UVicAPI) GetAllCourses() ([]byte, error) {
	res, err := http.Get(kuali)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(res.Body); err != nil {
		return nil, err
	}

	resJson, err := fastjson.ParseBytes(buf.Bytes())
	if err != nil {
		return nil, err
	}

	courses := []Course{}

	v := resJson.GetArray()
	for _, course := range v {
		id, err := course.Get("__catalogCourseId").StringBytes()
		if err != nil {
			return nil, err
		}

		title, err := course.Get("title").StringBytes()
		if err != nil {
			return nil, err
		}

		c := Course{
			ID:    string(id),
			Title: string(title),
		}

		courses = append(courses, c)
	}

	return json.MarshalIndent(courses, "", "  ")
}
