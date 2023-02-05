package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hn275/uvic-api/uvicapi"
)

func main() {

	term := "202301"

	log.Printf("fetching uvic data for term %s\n", term)
	uvic, err := uvicapi.NewAPI(term)
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
		}

		res, err := uvic.GetSection(q)
		if err != nil {
			log.Fatal(err)
		}

		write("GetSection", res)
	}

	// GetAllSections
	{
		res, err := uvic.GetAllSections(1)
		if err != nil {
			log.Fatal(err)
		}

		write("GetAllSections", res)
	}

	// GetCourseDesc
	{
		_, err := uvic.GetCourseDesc("20747") // wtf?
		if err != nil {
			log.Fatal(err)
		}
		// TODO: what
	}

	// GetTerms
	{
		res, err := uvic.GetTerms()

		if err != nil {
			panic(err)
		}

		write("GetTerms", res)
	}

	// GetAllCourses
	{
		res, err := uvic.GetAllCourses()
		if err != nil {
			log.Fatal(err)
		}

		write("GetAllSections", res)
	}
}

func write(fnName string, data []byte) {
	if err := os.WriteFile(fmt.Sprintf("./data/%s.json", fnName), data, 0666); err != nil {
		log.Printf("ERROR [%s]:\n%v\n", fnName, err)
	} else {
		log.Printf("DONE [%s]\n", fnName)
	}
}
