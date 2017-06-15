package afenav

type executeReportRequest struct {
	ReportQuery
	AuthenticationToken authenticationToken
}

// ReportQuery defines a request to the Document Reporting API
type ReportQuery struct {
	DocumentType    string
	ReportType      string
	Columns         []string
	SortColumns     []ReportSortColumn
	GlobalSearch    string
	Filter          []ReportFilter
	SkipRows        int
	MaxRowCount     int
	IncludeArchived bool
}

// ReportSortColumn defines a sorting column and direction
type ReportSortColumn struct {
	Column    string
	Ascending bool
}

// ReportFilter represents a report filter clause
type ReportFilter struct {
	LeftParenthesis  string
	Column           string
	Operator         string
	Value            string
	RightParenthesis string
	Join             string
}

// ExecuteReportResponse represents the response from the ExecuteReport API
type ExecuteReportResponse struct {
	TotalRowCount    int
	FilteredRowCount int
	ElapsedTimeMS    float64
	Columns          []ExecuteReportColumn
	Rows             []ExecuteReportRow
}

//ExecuteReportColumn represents column information returns from the ExecuteReport API
type ExecuteReportColumn struct {
	UniqueDisplayName string
	ColumnID          string `json:"ColumnId"`
	Name              string
	Type              string
}

// ExecuteReportRow represents a row of display friendly data from the ExecuteReport API
type ExecuteReportRow struct {
	DocumentID DocumentID `json:"DocumentId"`
	Data       []string
}

// ExecuteReport executes a document report and returns the formatted results
func (service *Service) ExecuteReport(query ReportQuery) (*ExecuteReportResponse, error) {
	var response ExecuteReportResponse

	if err := service.invokeJSON("/api/Documents/Reporting/Execute", executeReportRequest{
		ReportQuery:         query,
		AuthenticationToken: service.authenticationToken,
	}, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
