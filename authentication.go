package afenav

import "errors"

type LoginRequest struct {
	Username string `json:"UserName"`
	Password string
}

type LoginResponse struct {
	AuthenticationToken AuthenticationToken
}

// Login opens a session against the AFE Navigator service and stores the authenticationToken
func (service *Service) Login() error {
	var response LoginResponse
	if err := service.invokeJSON("/api/Authentication/Login", LoginRequest{
		Username: service.Config.Username,
		Password: service.Config.Password,
	}, &response); err != nil {
		return err
	}

	service.AuthenticationToken = response.AuthenticationToken

	return nil
}

// Logout terminates the active session and erases the authenticationToken
func (service *Service) Logout() error {
	if service.AuthenticationToken == "" {
		return errors.New("Not logged in")
	}
	if err := service.invokeJSON("/api/Authentication/Logout", AuthenticationTokenRequest{
		AuthenticationToken: service.AuthenticationToken,
	}, nil); err != nil {
		return err
	}

	service.AuthenticationToken = ""
	return nil
}
