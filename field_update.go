package afenav

import "time"

type UpdateTextRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          string
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
}

type UpdateUwiRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          string
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
}

type UpdateDateRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          time.Time
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
}

type UpdateIntegerRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          int
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
}

type UpdateBooleanRequest struct {
	FieldPath           []string
	FieldName           string
	FieldValue          bool
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
}

func (service *Service) UpdateText(documentHandle DocumentHandle, path []string, field string, value string) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateText", UpdateTextRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateUwi(documentHandle DocumentHandle, path []string, field string, value string) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateUwi", UpdateUwiRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateInteger(documentHandle DocumentHandle, path []string, field string, value int) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateInteger", UpdateIntegerRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateBoolean(documentHandle DocumentHandle, path []string, field string, value bool) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateBoolean", UpdateBooleanRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value,
	}, nil); err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateDate(documentHandle DocumentHandle, path []string, field string, value time.Time) error {
	if err := service.invokeJSON("/api/Documents/Field/UpdateDate", UpdateDateRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      documentHandle,
		FieldPath:           path,
		FieldName:           field,
		FieldValue:          value.UTC(),
	}, nil); err != nil {
		return err
	}
	return nil
}
