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
		}

		res, err := uvicClient.GetSection(q)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("./data/GetSection.json", res, 0666); err != nil {
			log.Printf("ERROR [GetSection]:\n%v\n", time.Since(startTime))
		} else {
			log.Printf("DONE [GetSection] in %v\n", time.Since(startTime))
		}
	}

	// GetAllSections
	{
		res, err := uvicClient.GetAllSections(1)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("./data/GetAllSections.json", res, 0666); err != nil {
			log.Printf("ERROR [GetAllSections]:\n%v\n", time.Since(startTime))
		} else {
			log.Printf("DONE [GetAllSections] in %v\n", time.Since(startTime))
		}
	}

	// GetCourseDesc
	{
		_, err := uvicClient.GetCourseDesc("20747") // wtf?
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("DONE [GetCourseDesc] in %v\n", time.Since(startTime))
	}

	// GetTerms
	{
		res, err := uvicClient.GetTerms()

		if err != nil {
			panic(err)
		}
		if err := os.WriteFile("./data/GetTerms.json", res, 0666); err != nil {
			log.Printf("ERROR [GetTerms]:\n%v", err)
		} else {
			log.Printf("DONE [GetTerms] in %v\n", time.Since(startTime))
		}
	}

	// GetAllCourses
	{
		res, err := uvicClient.GetAllCourses()
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile("./data/GetAllCourses.json", res, 0666); err != nil {
			log.Printf("ERROR [GetAllCourses]:\n%v", err)
		} else {
			log.Printf("DONE [GetAllCourses] in %v\n", time.Since(startTime))
		}
	}

	log.Fatalf("DONE in %v\n", time.Since(startTime))
}
