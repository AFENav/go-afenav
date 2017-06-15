package afenav

import (
	"errors"
	"time"
)

// UwiValue is a struct that stores the UWI display value, type and sorted representation
// as returned by the AFE Nav service
type UwiValue struct {
	Value     string
	SortedUwi string
	UwiType   string
}

// DocumentField represents a field and value from the document read call
type DocumentField struct {
	ID                 string `json:"ID"`
	Text               string
	Bool               bool
	NumberDecimal      float64
	NumberInteger      int
	Date               time.Time
	Uwi                UwiValue
	GUID               string `json:"GUID"`
	Document           DocumentID
	DocumentDescriptor string
	Record             DocumentRecord
	Records            []DocumentRecord
	FileSize           int64
	FileType           string
	FileVersionNumber  int32
}

// DocumentRecord represents a record on a document (top level, custom, or list)
type DocumentRecord struct {
	Fields []DocumentField
}

// Field returns the field of the given ID from the record
func (record DocumentRecord) Field(id string) (*DocumentField, error) {
	for _, f := range record.Fields {
		if f.ID == id {
			return &f, nil
		}
	}
	return nil, errors.New("Field Not Found")
}

// DocumentData represents a snapshot of a single document
type DocumentData struct {
	DocumentID   DocumentID `json:"DocumentId"`
	DocumentType string
	Record       DocumentRecord
}

// Field returns the top-level filed of the given ID
func (doc DocumentData) Field(id string) (*DocumentField, error) {
	return doc.Record.Field(id)
}

// DocumentReadResponse represents the response from a read call
type DocumentReadResponse struct {
	BaseDocument   DocumentData
	ChildDocuments []DocumentData
}

// ChildDocument finds a child document record by its DocumentID
func (r DocumentReadResponse) ChildDocument(id DocumentID) (*DocumentData, error) {
	for _, d := range r.ChildDocuments {
		if d.DocumentID == id {
			return &d, nil
		}
	}
	return nil, errors.New("Document not found")
}

type documentReadRequest struct {
	AuthenticationToken    authenticationToken
	DocumentHandle         string
	SerializeDocumentTypes []string
}

// ReadDocument reads the full contents of the provided document, and any references document types listed in serializeDocumentTypes
func (service *Service) ReadDocument(handle DocumentHandle, serializeDocumentTypes []string) (*DocumentReadResponse, error) {
	var response DocumentReadResponse
	if err := service.invokeJSON("/api/Documents/Read", documentReadRequest{
		AuthenticationToken:    service.AuthenticationToken,
		DocumentHandle:         string(handle),
		SerializeDocumentTypes: serializeDocumentTypes,
	}, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type readCurrentRequest struct {
	DocumentID             DocumentID `json:"DocumentId"`
	DocumentType           string
	SerializeDocumentTypes []string
	AuthenticationToken    authenticationToken
}

// ReadCurrent returns a snapshot of the current data for a document
func (service *Service) ReadCurrent(documentType string, documentID DocumentID, serializeDocumentTypes []string) (*DocumentReadResponse, error) {
	var response DocumentReadResponse
	if err := service.invokeJSON("/api/Documents/ReadCurrent", readCurrentRequest{
		AuthenticationToken:    service.AuthenticationToken,
		DocumentType:           documentType,
		DocumentID:             documentID,
		SerializeDocumentTypes: serializeDocumentTypes,
	}, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
