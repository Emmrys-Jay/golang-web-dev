package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type tField struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

type tableFields []tField

func main() {
	csvFile, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close()

	csvLine, err := csv.NewReader(csvFile).Read()
	if err != nil {
		log.Fatalln(err)
	}
	csvLine1, err := csv.NewReader(csvFile).Read()
	if err != nil {
		log.Fatalln(err)
	}

	tableField1 := tableFields{
		tField{
			csvLine[0],
			csvLine[1],
			csvLine[2],
			csvLine[3],
			csvLine[4],
			csvLine[5],
			csvLine[6],
		},

		tField{
			csvLine1[0],
			csvLine1[1],
			csvLine1[2],
			csvLine1[3],
			csvLine1[4],
			csvLine1[5],
			csvLine1[6],
		},
	}

	fmt.Println(tableField1)
}
