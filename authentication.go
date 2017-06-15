package afenav

import "errors"

type loginRequest struct {
	Username string `json:"UserName"`
	Password string
}

type loginResponse struct {
	AuthenticationToken authenticationToken
}

// Login opens a session against the AFE Navigator Service and stores the authenticationToken
func (service *Service) Login(username, password string) error {
	var response loginResponse
	if err := service.invokeJSON("/api/Authentication/Login", loginRequest{
		Username: username,
		Password: password,
	}, &response); err != nil {
		return err
	}

	service.authenticationToken = response.AuthenticationToken

	return nil
}

// Logout terminates the active session and erases the authenticationToken
func (service *Service) Logout() error {
	if service.authenticationToken == "" {
		return errors.New("Not logged in")
	}
	if err := service.invokeJSON("/api/Authentication/Logout", authenticationTokenRequest{
		AuthenticationToken: service.authenticationToken,
	}, nil); err != nil {
		return err
	}

	service.authenticationToken = ""
	return nil
}
