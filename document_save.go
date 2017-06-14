package afenav

// SaveDocument writes pending changes for a document to the database
func (service *Service) SaveDocument(handle DocumentHandle) error {
	if err := service.invokeJSON("/api/Documents/Save", documentHandleRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
