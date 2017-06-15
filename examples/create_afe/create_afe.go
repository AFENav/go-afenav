package main

import (
	"time"

	"path"

	"github.com/AFENav/go-afenav"
	"github.com/BurntSushi/toml"
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
	if _, err := toml.DecodeFile(path.Join("..", "service.config"), &config); err != nil {
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

	if err = service.SetPrimaryWell(documentHandle, listItemID); err != nil {
		panic(err)
	}

}
