package api

import (
	"errors"
	"net/http"
)

const (
	appControlPath = "/sony/appControl"
)

// AppControlService handles requests related to listing and opening apps
type AppControlService service

type GetApplicationListResult = Result[[1][]struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
	Icon  string `json:"icon"`
}]

type getApplicationListParams [0]struct{}
type getApplicationListPayload Payload[getApplicationListParams]

func (s *AppControlService) GetApplicationList() (*GetApplicationListResult, *http.Response, error) {
	body := getApplicationListPayload{
		Method:  "getApplicationList",
		ID:      1,
		Params:  getApplicationListParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, appControlPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetApplicationListResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

type SetActiveAppResult = Result[[0]struct{}]

type setActiveAppParams [1]struct {
	URI  string  `json:"uri"`
	Data *string `json:"data,omitempty"`
}
type setActiveAppPayload Payload[setActiveAppParams]

func (s *AppControlService) SetActiveApp(uri string, data *string) (*SetActiveAppResult, *http.Response, error) {
	body := setActiveAppPayload{
		Method:  "setActiveApp",
		ID:      1,
		Params:  setActiveAppParams{{URI: uri, Data: data}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, appControlPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(SetActiveAppResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}
