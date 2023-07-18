package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "12345678"

	// Create a new Neo4j driver
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Open a new session for each request
		session, err := driver.Session(neo4j.AccessModeRead)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to open session", http.StatusInternalServerError)
			return
		}
		defer session.Close()

		// Run a Cypher query to fetch data from the knowledge graph with relationships and labels
		query := "MATCH (n)-[r]->(m) RETURN n, labels(n) AS n_labels, r, type(r) AS r_type, m, labels(m) AS m_labels LIMIT 10"
		result, err := session.Run(query, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to execute query", http.StatusInternalServerError)
			return
		}

		// Process the query results
		var data []map[string]interface{}
		for result.Next() {
			record := result.Record()
			node := record.GetByIndex(0).(neo4j.Node)
			nodeLabels := record.GetByIndex(1).([]interface{})
			relationship := record.GetByIndex(2).(neo4j.Relationship)
			relationshipType := record.GetByIndex(3).(string)
			relatedNode := record.GetByIndex(4).(neo4j.Node)
			relatedNodeLabels := record.GetByIndex(5).([]interface{})

			nodeProps := node.Props
			relationshipProps := relationship.Props
			relatedNodeProps := relatedNode.Props

			data = append(data, map[string]interface{}{
				"node":              nodeProps,
				"node_labels":       nodeLabels,
				"relationship":      relationshipProps,
				"relationship_type": relationshipType,
				"related_node":      relatedNodeProps,
				"related_labels":    relatedNodeLabels,
			})
		}

		// Convert the data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to serialize data", http.StatusInternalServerError)
			return
		}

		// Set the JSON response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Write the JSON response
		_, err = w.Write(jsonData)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
