package afenav

type CreateRequest struct {
	AuthenticationToken AuthenticationToken
	DocumentType        string
	DefaultValues       []CreateRequestDefaultValue
}

type CreateRequestDefaultValue struct {
	Path  string
	Value string
}

type CreateResponse struct {
	DocumentID     DocumentID `json:"DocumentId"`
	DocumentHandle DocumentHandle
}

func (service *Service) Create(documentType string) (*CreateResponse, error) {
	var response CreateResponse
	if err := service.invokeJSON("/api/Documents/Create", CreateRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentType:        documentType,
	}, &response); err != nil {
		return nil, err
	}
	service.Log.Infof("Created new %v document with ID '%v'", documentType, response.DocumentID)
	return &response, nil
}
