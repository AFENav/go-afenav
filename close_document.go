package afenav

// Close a document handle!
func (service *Service) CloseDocument(handle DocumentHandle) error {
	if err := service.invokeJSON("/api/Documents/Close", DocumentHandleRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
