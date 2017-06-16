package afenav

// CloseDocument releases the lock on a document and saves any pending edits
func (service *Service) CloseDocument(handle DocumentHandle) error {
	if err := service.invoke("/api/Documents/Close", documentHandleRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      string(handle),
	}, nil); err != nil {
		return err
	}
	return nil
}
