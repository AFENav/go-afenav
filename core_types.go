package afenav

// DocumentID represents the unique ID for a document (a GUID)
type DocumentID string

// AuthenticationToken is a string that represents a user session
type AuthenticationToken string

// DocumentHandle represents an open handle to a document in AFE Nav
type DocumentHandle string

// ListItemID represents a record in a document list
type ListItemID string

// DocumentHandleRequest represents a basic request for an in-memory document
type DocumentHandleRequest struct {
	AuthenticationToken AuthenticationToken
	DocumentHandle      string
}

// AuthenticationTokenRequest represents a basic authenticated request
type AuthenticationTokenRequest struct {
	AuthenticationToken AuthenticationToken
}

// Error is the standard format for unexpected exceptions from AFE Nav
type Error struct {
	ClassName string
	Message   string
}

// Config represents the configuration for an AFE Nav with user credentials
type Config struct {
	URL                string
	Username           string
	Password           string
	InsecureSkipVerify bool
	LogRequests        bool
}
