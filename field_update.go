package afenav

import "time"

type updateTextRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          string
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type updateDocumentReferenceRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          DocumentID
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type updateUwiRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          string
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type updateDateRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          time.Time
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type updateIntegerRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          int
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type updateDecimalRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          float64
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type updateBooleanRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          bool
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

// UpdateText sets the value for a text field on a document
func (service *Service) UpdateText(documentHandle DocumentHandle, path []string, field string, value string) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateText", updateTextRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

// UpdateDocumentReference updates the value for a document field on a document
func (service *Service) UpdateDocumentReference(documentHandle DocumentHandle, path []string, field string, value DocumentID) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateDocumentReference", updateDocumentReferenceRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

// UpdateUWI updates the value for a UWI type field on a document
func (service *Service) UpdateUWI(documentHandle DocumentHandle, path []string, field string, value string) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateUwi", updateUwiRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

// UpdateInteger updates the value for an integer field on a document
func (service *Service) UpdateInteger(documentHandle DocumentHandle, path []string, field string, value int) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateInteger", updateIntegerRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

// UpdateDecimal updates the value for a decimal field on a document
func (service *Service) UpdateDecimal(documentHandle DocumentHandle, path []string, field string, value float64) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateDecimal", updateDecimalRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

// UpdateBoolean updates the value for a boolean field for a document
func (service *Service) UpdateBoolean(documentHandle DocumentHandle, path []string, field string, value bool) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateBoolean", updateBooleanRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

// UpdateDate updates the value for a date/time field for a document
func (service *Service) UpdateDate(documentHandle DocumentHandle, path []string, field string, value time.Time) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateDate", updateDateRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value.UTC(),
	}, nil); err != nil {
		return err
	}
	return nil
}

type createRecordListItemRequest struct {
	FieldPath           []string
	FieldName           string
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

type createRecordListItemResponse struct {
	RecordListItemIdentifier ListItemID
}

// AddListItem adds a new item to a record list field
func (service *Service) AddListItem(handle DocumentHandle, path []string, field string) (ListItemID, error) {
	var response createRecordListItemResponse
	if err := service.invokeJSON("/api/Documents/Field/CreateRecordListItem", createRecordListItemRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      handle,
		FieldName:           field,
		FieldPath:           path,
	}, &response); err != nil {
		return ListItemID(""), err
	}
	return response.RecordListItemIdentifier, nil
}

type removeRecordListItemRequest struct {
	FieldPath           []string
	FieldName           string
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
}

// RemoveListItem adds a new item to a record list field
func (service *Service) RemoveListItem(handle DocumentHandle, path []string, field string, listItem ListItemID) error {
	if err := service.invokeJSON("/api/Documents/Field/RemoveRecordListItem", createRecordListItemRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      handle,
		FieldName:           field,
		FieldPath:           path,
	}, nil); err != nil {
		return err
	}
	return nil
}
