package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func connectToNeo4j() (neo4j.Driver, neo4j.Session, error) {
	uri := "bolt://localhost:7687" // Replace with your Neo4j database URI
	username := "neo4j"            // Replace with your Neo4j database username
	password := "12345678"         // Replace with your Neo4j database password

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, nil, err
	}

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, nil, err
	}

	return driver, session, nil
}

func loadFileToNeo4j(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	driver, session, err := connectToNeo4j()
	if err != nil {
		return err
	}
	defer driver.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Process each line and extract the relevant information
		// Example: Assuming the file format is "<node1>,<node2>"
		nodes := strings.Split(line, ",")
		node1 := nodes[0]
		node2 := nodes[1]

		// Create a Cypher query to create the nodes and relationship
		cypherQuery := `
            MERGE (n1:Node {id: $node1})
            MERGE (n2:Node {id: $node2})
            MERGE (n1)-[:CONNECTED_TO]->(n2)
        `

		_, err := session.Run(cypherQuery, map[string]interface{}{
			"node1": node1,
			"node2": node2,
		})
		if err != nil {
			log.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	filePath := "/path/to/your/file.txt" // Replace with the path to your file
	err := loadFileToNeo4j(filePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File loaded to Neo4j successfully!")
}
