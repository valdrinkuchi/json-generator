package engine

import (
	"encoding/json"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Answer struct {
	Key            string            `json:"key"`
	Active         bool              `json:"active"`
	Label          map[string]string `json:"label"`
	ReportingLabel map[string]string `json:"reporting_label"`
}

func Generator() {
	object := &Answer{}
	var result []Answer
	data, err := excelize.OpenFile("answers_negative.xlsx")
	if err != nil {
		log.Fatalf("Error occured during reading the Excel file. Error: %s", err.Error())
	}
	sheetName := data.GetSheetName(0)
	rowData, err := data.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Error occured during reading rows. Error: %s", err.Error())
	}
	for _, row := range rowData {
		if row[0] == "key" {
			continue
		}
		object.Key = row[0]
		object.Active = true
		label := make(map[string]string)
		for i := 1; i < len(row); i++ {
			label[rowData[0][i]] = row[i]
		}
		object.Label = label
		object.ReportingLabel = object.Label
		result = append(result, *object)
	}
	resultJSON, err := json.MarshalIndent(result, "", " ")
	f, err := os.Create("test.json")
	if err != nil {
		log.Fatalf("Could not create the file. Error: %s", err.Error())
	}
	f.WriteString(string(resultJSON))
	if err != nil {
		log.Fatalf("Could not write into the file. Error: %s", err.Error())
		f.Close()
	}
}
