package afenav

import "errors"

type findDocumentsRequest struct {
	DocumentType string
	AuthenticationToken authenticationToken
	SearchStrings []string
}

type findDocumentsResponse struct {
	FoundDocuments []foundDocument
}

type foundDocument struct {
	SearchString string
	DocumentID DocumentID `json:"MatchingDocument"`
}

// FindDocument finds a document by a search string and returns the ID, or an error if it doesn't match a single document
func (service *Service) FindDocument(documentType string, searchString string) (DocumentID, error) {
	var response findDocumentsResponse
	if err := service.invoke("/api/Documents/FindDocuments", findDocumentsRequest{
		AuthenticationToken: service.authenticationToken,
		DocumentType:        documentType,
		SearchStrings:       []string {searchString},
	}, &response); err != nil {
		return "", err
	}

	for _, tmp := range response.FoundDocuments {
		if tmp.SearchString == searchString {
			return tmp.DocumentID, nil
		}
	}

	return "", errors.New("No matching document found")
}