package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type TableData struct {
	Rank           string
	Country        string
	Area           string
	Water string
}

func writeToCSV(data []TableData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Rank", "Country", "Area", "Water"}
	writer.Write(headers)

	for _, record := range data {
		row := []string{record.Rank, record.Country, record.Area, record.Water}
		writer.Write(row)
	}

	return nil
}

func scrapeTable(url string, tableIndex int) ([]TableData, error) {
	var tableData []TableData

	c := colly.NewCollector()
	currentTableIndex := 0

	c.OnHTML(".wikitable > tbody", func(h *colly.HTMLElement) {
		currentTableIndex++
		if currentTableIndex == tableIndex {
			rowIndex := 0
			h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				if rowIndex == 0 {
					rowIndex++
					return
				}
				country := el.ChildText("td:nth-child(2)")
				area := el.ChildText("td:nth-child(3)")
				water := el.ChildText("td:nth-child(4)")
				if water == "" {
					water = "N/A"
				}
				data := TableData{
					Rank:           el.ChildText("td:nth-child(1)"),
					Country:        country,
					Area:           area,
					Water: strings.TrimSpace(water),
				}
				tableData = append(tableData, data)
			})
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return tableData, nil
}

func main() {
	url := "https://en.wikipedia.org/wiki/List_of_countries_and_dependencies_by_area"
	tableIndex := 1
	outputFile := "largest_countries_by_area.csv"

	tableData, err := scrapeTable(url, tableIndex)
	if err != nil {
		fmt.Println("Error scraping table:", err)
		return
	}

	err = writeToCSV(tableData, outputFile)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	fmt.Println("Data has been written to", outputFile)
}
