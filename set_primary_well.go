package afenav

type SetPrimaryWellRequest struct {
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
	ListItemID          ListItemID `json:"ListItemID"`
}

func (service *Service) SetPrimaryWell(documentHandle DocumentHandle, listItemID ListItemID) error {
	if err := service.invokeJSON("/api/Afe/SetPrimaryWell", SetPrimaryWellRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      documentHandle,
		ListItemID:          listItemID,
	}, nil); err != nil {
		return err
	}
	return nil
}
