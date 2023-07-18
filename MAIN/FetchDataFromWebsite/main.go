package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type University struct {
	Name                string `json:"name"`
	UnderGraduateCourse string `json:"ugcourse"`
	PostGraduateCourse  string `json:"pgcourse"`
	Website             string `json:"link"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("University_Batch_1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV file
	var universities []University

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		university := University{
			Name:                record[0],
			UnderGraduateCourse: record[1],
			PostGraduateCourse:  record[2],
			Website:             record[3],
		}

		universities = append(universities, university)
	}

	// Convert universities slice to JSON
	jsonData, err := json.MarshalIndent(universities, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Save JSON data to a file
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conversion completed successfully.")
}
