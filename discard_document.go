package afenav

// Close a document handle!
func (service *Service) DiscardDocument(handle DocumentHandle) error {
	if err := service.invokeJSON("/api/Documents/Discard", DocumentHandleRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
