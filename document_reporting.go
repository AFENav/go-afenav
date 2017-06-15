package afenav

type documentSummaryRequest struct {
	AuthenticationToken      authenticationToken
	Type                     string
	ColumnIds                []DocumentSummaryColumn
	BaseFilter               string
	TopLevelSearch           string
	ColumnSorts              []DocumentSummaryColumnSort
	RowStartIndex            int32
	Count                    int32
	UnformattedNumericValues bool
	Echo                     int32
}

// DocumentSummaryColumn represents a column to include in browse results
type DocumentSummaryColumn struct {
	ColumnID string `json:"Id"`
	Filter   string
}

// DocumentSummaryColumnSort represents a column to sort on
type DocumentSummaryColumnSort struct {
	ColumnID      string `json:"Id"`
	SortAscending bool
}

// DocumentSummaryResponse represents a response from the reporting API
type DocumentSummaryResponse struct {
	Rows []DocumentSummaryRow
}

// DocumentSummaryRow represents a single row from the reporting API
type DocumentSummaryRow struct {
	DocumentID DocumentID `json:"Id"`
}

// ListDocuments lists all DocumentIDs for Documents of a specific DocumentType
func (service *Service) ListDocuments(documentType string) ([]DocumentID, error) {
	var response DocumentSummaryResponse
	if err := service.invokeJSON("/api/Documents/Summary", documentSummaryRequest{
		AuthenticationToken: service.authenticationToken,
		Type:                documentType,
		BaseFilter:          "All",
		RowStartIndex:       0,
		Echo:                1,
		ColumnSorts:         []DocumentSummaryColumnSort{},
		ColumnIds:           []DocumentSummaryColumn{},
		Count:               999999,
	}, &response); err != nil {
		return nil, err
	}

	ids := make([]DocumentID, len(response.Rows))
	for i := 0; i < len(response.Rows); i++ {
		ids[i] = response.Rows[i].DocumentID
	}
	return ids, nil
}
