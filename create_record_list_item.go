package afenav

type CreateRecordListItemRequest struct {
	FieldPath           []string
	FieldName           string
	DocumentHandle      DocumentHandle
	AuthenticationToken AuthenticationToken
}

type CreateRecordListItemResponse struct {
	RecordListItemIdentifier ListItemID
}

func (service *Service) AddListItem(handle DocumentHandle, path []string, field string) (ListItemID, error) {
	var response CreateRecordListItemResponse
	if err := service.invokeJSON("/api/Documents/Field/CreateRecordListItem", CreateRecordListItemRequest{
		AuthenticationToken: service.AuthenticationToken,
		DocumentHandle:      handle,
		FieldName:           field,
		FieldPath:           path,
	}, &response); err != nil {
		return ListItemID(""), err
	}
	return response.RecordListItemIdentifier, nil
}
