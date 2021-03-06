package afenav

type setPrimaryWellRequest struct {
	DocumentHandle      DocumentHandle
	AuthenticationToken authenticationToken
	ListItemID          ListItemID `json:"ListItemID"`
}

// SetPrimaryWell will mark a well list item as the primary well.  The record must already exist.
func (service *Service) SetPrimaryWell(documentHandle DocumentHandle, listItemID ListItemID) error {
	if err := service.invoke("/api/Afe/SetPrimaryWell", setPrimaryWellRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentHandle:      documentHandle,
		ListItemID:          listItemID,
	}, nil); err != nil {
		return err
	}
	return nil
}
