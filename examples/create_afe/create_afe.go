package main

import (
	"time"

	"github.com/AFENav/go-afenav"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

// =============================================================
// AFE Nav API sample:
// 1) login to API
// 2) create a sample AFE
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

	// create a new AFE, returning the document handle
	_, documentHandle, err := service.CreateDocument("AFE")
	if err != nil {
		panic(err)
	}

	// always close document handles to avoid keeping locks on them
	defer service.CloseDocument(documentHandle)


	// update AFE Type PAA by searching for DocumentID for "DRILLING" and then setting CUSTOM/AFE_TYPE
	var drillingDocumentID afenav.DocumentID
	if drillingDocumentID, err = service.FindDocument("LUT_AFE_TYPE", "DRILLING"); err != nil {
		panic(err)
	}
	if err = service.UpdateDocumentReference(documentHandle, []string{"CUSTOM"}, "AFE_TYPE", drillingDocumentID); err != nil {
		panic(err)
	}

	// update Operator Status PAA by searching for DocumentID for "DRILLING" and then setting CUSTOM/AFE_TYPE
	var operatedStatus afenav.DocumentID
	if operatedStatus, err = service.FindDocument("LUT_OPRTR_STTS", "Operated"); err != nil {
		panic(err)
	}
	if err = service.UpdateDocumentReference(documentHandle, []string{"CUSTOM"}, "OPERATOR_STATUS", operatedStatus); err != nil {
		panic(err)
	}

	if err = service.UpdateText(documentHandle, []string{}, "DESCRIPTION", "Hello World"); err != nil {
		panic(err)
	}

	if err = service.UpdateDate(documentHandle, []string{}, "START_DATE", time.Now()); err != nil {
		panic(err)
	}

	if err = service.UpdateInteger(documentHandle, []string{}, "BUDGET_YEAR", 2017); err != nil {
		panic(err)
	}

	if err = service.UpdateBoolean(documentHandle, []string{"CUSTOM"}, "BUDGETED", true); err != nil {
		panic(err)
	}

	var listItemID afenav.ListItemID
	if listItemID, err = service.AddListItem(documentHandle, []string{}, "WELL"); err != nil {
		panic(err)
	}

	if err = service.UpdateUWI(documentHandle, []string{"WELL", string(listItemID)}, "UWI", "81278126378678"); err != nil {
		panic(err)
	}
	if err = service.UpdateDecimal(documentHandle, []string{"WELL", string(listItemID)}, "UWI_LAT", -114.1234); err != nil {
		panic(err)
	}
	if err = service.UpdateDecimal(documentHandle, []string{"WELL", string(listItemID)}, "UWI_LONG", 100.123); err != nil {
		panic(err)
	}

	if err = service.SetPrimaryWell(documentHandle, listItemID); err != nil {
		panic(err)
	}

}
