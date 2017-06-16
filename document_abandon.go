package afenav

// AbandonDocument discards pending changes to a document handle!
func (service *Service) AbandonDocument(handle DocumentHandle) error {
	if err := service.invoke("/api/Documents/AbandonChanges", documentHandleRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
