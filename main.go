package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Record struct {
	Month   string
	YearI   int
	YearII  int
	YearIII int
}

func main() {
	// Open the CSV file
	file, err := os.Open("airtravel.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader with headers
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	headers, err := reader.Read()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Map headers to column indices
	monthIdx := -1
	year1Idx := -1
	year2Idx := -1
	year3Idx := -1

	for i, header := range headers {
		switch header {
		case "Month":
			monthIdx = i
		case "1947":
			year1Idx = i
		case "1957":
			year2Idx = i
		case "1967":
			year3Idx = i
		}
	}

	//checking wheather headers index is assigned or not
	if monthIdx == -1 || year1Idx == -1 || year2Idx == -1 || year3Idx == -1 {
		fmt.Println("Error: missing headers")
		return
	}

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a slice to hold the parsed records
	var parsedRecords []Record

	// Parse each record and append it to the slice
	for _, record := range records {
		month := record[monthIdx]
		yearI := record[year1Idx]
		yearII := record[year2Idx]
		yearIII := record[year3Idx]

		// Convert the age from string to int
		year1, err1 := strconv.Atoi(yearI)
		year2, err2 := strconv.Atoi(yearII)
		year3, err3 := strconv.Atoi(yearIII)

		if err1 != nil || err2 != nil || err3 != nil {

			fmt.Println("Error1: ", err1)
			fmt.Println("Error2: ", err2)
			fmt.Println("Error3: ", err3)

			return
		}

		parsedRecords = append(parsedRecords, Record{
			Month:   month,
			YearI:   year1,
			YearII:  year2,
			YearIII: year3,
		})
	}

	// Output the parsed records as a table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, record := range parsedRecords {
		table.Append([]string{record.Month, strconv.Itoa(record.YearI), strconv.Itoa(record.YearII), strconv.Itoa(record.YearIII)})
	}
	table.Render()
}
