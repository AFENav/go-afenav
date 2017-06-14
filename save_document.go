package afenav

// Close a document handle!
func (service *Service) SaveDocument(handle DocumentHandle) error {
	if err := service.invokeJSON("/api/Documents/Save", DocumentHandleRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
