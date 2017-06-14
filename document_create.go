package afenav

type createRequest struct {
	AuthenticationToken authenticationToken
	DocumentType        string
}

type createResponse struct {
	DocumentID     DocumentID `json:"DocumentId"`
	DocumentHandle DocumentHandle
}

// CreateDocument creates a new document of a specified type and returns the Document ID and Document Handle
func (service *Service) CreateDocument(documentType string) (DocumentID, DocumentHandle, error) {
	var response createResponse
	if err := service.invokeJSON("/api/Documents/Create", createRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentType:        documentType,
	}, &response); err != nil {
		return DocumentID(""), DocumentHandle(""), err
	}
	service.log.Infof("Created new %v document with ID '%v'", documentType, response.DocumentID)
	return response.DocumentID, response.DocumentHandle, nil
}
