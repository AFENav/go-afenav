package main

import (
	"path/filepath"

	"fmt"
	"github.com/AFENav/go-afenav"
	"github.com/BurntSushi/toml"
)

// =============================================================
// AFE Nav API sample:
// 1) login to API
// 2) report on AFEs
// 3) logout
// =============================================================

// Config represents the configuration for an AFE Nav with user credentials
type Config struct {
	URL                string
	Username           string
	Password           string
	InsecureSkipVerify bool
	LogRequests        bool
}

func main() {
	var config Config
	if _, err := toml.DecodeFile(filepath.Join("..", "service.config"), &config); err != nil {
		panic(err)
	}

	service := afenav.New(config.URL)
	service.LogRequests = config.LogRequests
	service.InsecureSkipVerify = config.InsecureSkipVerify

	// login.
	if err := service.Login(config.Username, config.Password); err != nil {
		panic(err)
	}

	// always logout or session will be left open until automatically suspended
	defer service.Logout()

	query := afenav.ReportQuery{
		DocumentType: "AFE",
		ReportType:   "AFE",
		Columns: []string{
			// "DOCUMENT_ID",
			"AFENUMBER_DOC/AFENUMBER",
			"STATUS",
			"TOTAL_GROSS_ESTIMATE",
		},
		SortColumns: []afenav.ReportSortColumn{
			{
				Column:    "TOTAL_GROSS_ESTIMATE",
				Ascending: false,
			},
		},
		Filter: []afenav.ReportFilter{
			{
				Column:   "STATUS",
				Operator: "<>",
				Value:    "UNREL",
				Join:     "AND",
			},
			{
				Column:   "AFENUMBER_DOC/AFENUMBER",
				Operator: "<>",
				Value:    "",
				Join:     "AND",
			},
		},
		SkipRows:        0,
		MaxRowCount:     9999,
		IncludeArchived: false,
	}

	data, err := service.ExecuteReport(query)
	if err != nil {
		panic(err)
	}

	for _, col := range data.Columns {
		fmt.Printf("%30v", col.UniqueDisplayName)
	}
	fmt.Printf("\n")

	for _, row := range data.Rows {
		for idx := range data.Columns {
			fmt.Printf("%30v", row.Data[idx])
		}
		fmt.Printf("\n")
	}
}
