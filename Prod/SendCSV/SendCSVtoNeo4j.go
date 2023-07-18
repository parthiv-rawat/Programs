package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	URI        = "bolt://localhost:7687"
	USERNAME   = "neo4j"
	PWD        = "12345678"
	importPath = "C:\\Users\\Admin\\.Neo4jDesktop\\relate-data\\dbmss\\dbms-46d21932-5afc-4bd8-87d8-f6054473e030\\import\\Programs\\"
)

func main() {
	uri := URI
	username := USERNAME
	password := PWD

	filePathForUpload := "file:///C:\\Programs\\University.csv"
	filePathForCoping := "C:\\Programs\\University.csv"

	filename := filepath.Base(filePathForUpload)

	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]

	parts := strings.Split(name, "_")
	firstFilename := parts[0]

	SearchList := []string{"School", "University", "Course", "Student", "Teacher", "UG", "PG", "Node"} // Add your list of filenames here

	if !contains(SearchList, firstFilename) {
		log.Printf("'%s' not found in the search list", firstFilename)
	}

	copyCSVfileToNeo4j(filePathForCoping, filename)

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Println(err)
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	query := `
	LOAD CSV WITH HEADERS FROM $filePath AS row
    CREATE (n:` + firstFilename + `)
    SET n += row
`

	// query := `
	// LOAD CSV WITH HEADERS FROM $filePath AS row
	// CREATE (n)
	// SET n += row
	// WITH n, row.Labels AS labels
	// UNWIND labels AS label
	// SET n :` + "`" + `{{label}}` + "`" + `
	// `

	params := map[string]interface{}{
		"filePath": filePathForUpload,
	}

	result, err := session.Run(query, params)
	if err != nil {
		log.Println(err)
	}

	if result.Err() != nil {
		log.Println(result.Err())
	}

	fmt.Println("CSV file imported successfully!")
}

func contains(list []string, item string) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}

func copyCSVfileToNeo4j(source, filename string) {

	destination := importPath + filename

	// Open the source file
	src, err := os.Open(source)
	if err != nil {
		fmt.Println("Failed to open source file:", err)
		return
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(destination)
	if err != nil {
		fmt.Println("Failed to create destination file:", err)
		return
	}
	defer dst.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("Failed to copy file:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
