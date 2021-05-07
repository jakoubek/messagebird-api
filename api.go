package messagebird_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	EndpointBaseUrl = "https://rest.messagebird.com"
)

type Api struct {
	apiKey string
}

func NewApi(apiKey string) *Api {
	return &Api{
		apiKey: apiKey,
	}
}

func (api *Api) getAuthorizationHeader() string {
	return fmt.Sprintf("AccessKey %s", api.apiKey)
}

func (api *Api) getEndpointUrl(service string) (string, error) {
	switch service {
	case "SMS":
		return fmt.Sprintf("%s/%s", EndpointBaseUrl, "messages"), nil
	}
	return "", errors.New(fmt.Sprintf("No valid service given: %s", service))
}

func (api *Api) SendSms(sms *Sms) (ServiceAnswer, error) {

	var answer ServiceAnswer

	reqBody, err := json.Marshal(sms)
	if err != nil {
		return answer, err
	}

	apiEndpoint, err := api.getEndpointUrl("SMS")
	if err != nil {
		return answer, err
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return answer, err
	}

	req.Header.Set("Authorization", api.getAuthorizationHeader())
	resp, err := client.Do(req)
	if err != nil {
		return answer, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return answer, err
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Print(body)
		return answer, errors.New(fmt.Sprintf("Error sending SMS - Status code: %d", resp.StatusCode))
	}

	err = json.Unmarshal(body, &answer)
	if err != nil {
		return answer, err
	}

	return answer, nil

}