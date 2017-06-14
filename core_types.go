package afenav

// DocumentID represents the unique ID for a document (a GUID)
type DocumentID string

// authenticationToken is a string that represents a user session
type authenticationToken string

// DocumentHandle represents an open handle to a document in AFE Nav
type DocumentHandle string

// ListItemID represents a record in a document list
type ListItemID string

// documentHandleRequest represents a basic request for an in-memory document
type documentHandleRequest struct {
	AuthenticationToken authenticationToken
	DocumentHandle      string
}

// authenticationTokenRequest represents a basic authenticated request
type authenticationTokenRequest struct {
	AuthenticationToken authenticationToken
}

// serviceError is the standard format for unexpected exceptions from AFE Nav
type serviceError struct {
	ClassName string
	Message   string
}
