package afenav

// AbandonDocument discards pending changes to a document handle!
func (service *Service) AbandonDocument(handle DocumentHandle) error {
	if err := service.invokeJSON("/api/Documents/AbandonChanges", documentHandleRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
