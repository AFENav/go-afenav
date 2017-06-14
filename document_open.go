package afenav

type searchAndOpenReadonlyRequest struct {
	DocumentType        string
	SearchString        string
	AuthenticationToken authenticationToken
}

type searchAndOpenRequest struct {
	DocumentType        string
	SearchString        string
	AuthenticationToken authenticationToken
	AutoCommit          bool
}
type openResponse struct {
	DocumentHandle DocumentHandle
}

// SearchAndOpenReadonly searches for and opens a readonly handle to a document of a given type
func (service *Service) SearchAndOpenReadonly(documentType string, searchString string) (DocumentHandle, error) {
	var response openResponse
	if err := service.invokeJSON("/api/Documents/SearchAndOpenReadonly", searchAndOpenReadonlyRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentType:        documentType,
		SearchString:        searchString,
	}, &response); err != nil {
		return "", err
	}

	return response.DocumentHandle, nil
}

// SearchAndOpen searches for and opens a read/write handle to a document of a given type
func (service *Service) SearchAndOpen(documentType string, searchString string) (DocumentHandle, error) {
	var response openResponse
	if err := service.invokeJSON("/api/Documents/SearchAndOpen", searchAndOpenRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentType:        documentType,
		SearchString:        searchString,
	}, &response); err != nil {
		return "", err
	}

	return response.DocumentHandle, nil
}
