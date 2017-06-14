package afenav

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"fmt"

	"github.com/op/go-logging"
)

// Service represents an instance of an AFE Navigator Service with configuration and state
type Service struct {
	url string
	log *logging.Logger

	InsecureSkipVerify  bool
	AuthenticationToken authenticationToken
	LogRequests         bool
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
	requestJSON, err := json.Marshal(request)
	if err != nil {
		service.log.Errorf("Failure invoking %v: %v", api, err)
		return err
	}

	responseBytes, err := service.invoke(api, requestJSON)
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

	response := responseBuffer.Bytes()

	if service.LogRequests {
		detailMessage := new(bytes.Buffer)

		detailMessage.WriteString("POST: " + api + "\n")
		detailMessage.WriteString("\nHeaders:\n")
		for headerName, headerValue := range req.Header {
			detailMessage.WriteString(fmt.Sprintf("  %v: %v\n", headerName, headerValue))
		}
		detailMessage.WriteString("\nRequest:\n")
		if err = writeJSON(detailMessage, request); err != nil {
			detailMessage.Write(request)
		}
		detailMessage.WriteString("\n\nResponse:\n")
		if err = writeJSON(detailMessage, response); err != nil {
			detailMessage.Write(response)
		}

		service.log.Debug(detailMessage)
	}
	return response, nil
}

func writeJSON(writer *bytes.Buffer, data []byte) error {
	buf := new(bytes.Buffer)
	if err := json.Indent(buf, data, "", "  "); err != nil {
		return err
	}

	writer.Write(buf.Bytes())

	return nil
}
