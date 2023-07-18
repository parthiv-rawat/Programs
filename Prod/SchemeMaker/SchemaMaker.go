package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Adapter interface {
	LoadKnowledgeGraph() error
	LoadCSVData() error
	ProcessData()
	Run()
}

type EducationAdapter struct {
	knowledgeGraphFile string
	csvFile            string
	knowledgeGraph     map[string]interface{}
	csvData            [][]string
}

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type Edge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func NewEducationAdapter(knowledgeGraphFile, csvFile string) Adapter {
	return &EducationAdapter{
		knowledgeGraphFile: knowledgeGraphFile,
		csvFile:            csvFile,
	}
}

func (ea *EducationAdapter) LoadKnowledgeGraph() error {
	data, err := os.ReadFile(ea.knowledgeGraphFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &ea.knowledgeGraph)
	if err != nil {
		return err
	}

	return nil
}

func (ea *EducationAdapter) LoadCSVData() error {
	file, err := ioutil.ReadFile(ea.csvFile)
	if err != nil {
		return err
	}

	r := csv.NewReader(strings.NewReader(string(file)))
	ea.csvData, err = r.ReadAll()
	if err != nil {
		return err
	}

	return nil
}

func (ea *EducationAdapter) ProcessData() {
	// Process the knowledge graph and CSV data as needed
	if ea.knowledgeGraph == nil {
		panic("Knowledge graph not loaded.")
	}

	conversionFromJSONtoCSV(ea.knowledgeGraph)

	if ea.csvData == nil {
		panic("CSV data not loaded.")
	}

}

func (ea *EducationAdapter) Run() {
	err := ea.LoadKnowledgeGraph()
	if err != nil {
		panic(err)
	}

	err = ea.LoadCSVData()
	if err != nil {
		panic(err)
	}

	ea.ProcessData()
}

func main() {
	knowledgeGraphFile := "knowledge_graph.json"
	csvFile := "EducationData.csv"

	adapter := NewEducationAdapter(knowledgeGraphFile, csvFile)
	adapter.Run()
}

func conversionFromJSONtoCSV(knowledgeGraph map[string]interface{}) {

	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Label"}
	writer.Write(header)

	for _, node := range knowledgeGraph["nodes"].([]interface{}) {
		nodeMap := node.(map[string]interface{})
		id := nodeMap["id"].(string)
		label := nodeMap["label"].(string)

		writer.Write([]string{id, label})
	}

	if err := writer.Error(); err != nil {
		panic(err)

	}
}
