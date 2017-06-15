// Package afenav contains wrapper methods for integrating with the AFE Navigator 2017 APIs
package afenav

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/op/go-logging"
)

// Service represents an instance of an AFE Navigator Service with configuration and state
type Service struct {
	url                 string
	log                 *logging.Logger
	authenticationToken authenticationToken

	InsecureSkipVerify bool
	LogRequests        bool
}

// New returns a new instance of an AFE Navigator Service
func New(url string) *Service {
	return &Service{
		url: url,
		log: logging.MustGetLogger("afenav"),
	}
}

// ---------------------------------- API HELPER METHODS -------------------------------------------

// invokeJSON calls an JSON API marshalling the request object, and unmarshalling into the response object
// response will be nil of error != nil
func (service *Service) invokeJSON(api string, request interface{}, response interface{}) error {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		service.log.Errorf("Failure invoking %v: %v", api, err)
		return err
	}

	responseBytes, err := service.invoke(api, requestBytes)
	if err != nil {
		service.log.Errorf("Failure invoking %v: %v", api, err)
		return err
	}

	if response != nil {
		if err := json.Unmarshal(responseBytes, &response); err != nil {
			service.log.Errorf("Failure invoking %v: %v", api, err)
			return err
		}
	}

	service.log.Debugf("Successfully invoked %v", api)

	if service.LogRequests {
		detailMessage := new(bytes.Buffer)

		detailMessage.WriteString("POST: " + api + "\n")
		detailMessage.WriteString("\nRequest:\n")
		json.Indent(detailMessage, requestBytes, "", " ")
		detailMessage.WriteString("\n\nResponse:\n")
		json.Indent(detailMessage, responseBytes, "", " ")

		service.log.Debug(detailMessage)
	}

	return nil
}

func (service *Service) invoke(api string, request []byte) ([]byte, error) {

	// TLS configuration to bypass TLS check if we are using a self-signed cert
	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: service.InsecureSkipVerify,
	}

	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		TLSClientConfig: tlsClientConfig,
	}

	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("POST", service.url+api, bytes.NewReader(request))

	// indicate to AFE Nav Service that we're calling the JSON APIs (as opposed to XML)
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusInternalServerError {
		// If we get a 500, decode the result and parse into an serviceError object
		var serviceError serviceError
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&serviceError)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(serviceError.Message)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Invalid API")
	}

	responseBuffer := new(bytes.Buffer)
	responseBuffer.ReadFrom(resp.Body)

	return responseBuffer.Bytes(), nil

}
