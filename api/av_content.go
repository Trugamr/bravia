package api

import (
	"errors"
	"net/http"
)

const (
	avContentPath = "/sony/avContent"
)

type ExternalInputStatus struct {
	URI    string `json:"uri"`
	Title  string `json:"title"`
	Label  string `json:"label"`
	Icon   string `json:"icon"`
	Status bool   `json:"status"`
}

// AVContentService handles requests related to AV content, such as setting inputs and playing content
type AVContentService service

// GetCurrentExternalInputsStatusResult is the response from the getCurrentExternalInputsStatus method
type GetCurrentExternalInputsStatusResult = Result[[1][]ExternalInputStatus]

type getCurrentExternalInputsStatusParams [0]struct{}
type getCurrentExternalInputsStatusPayload Payload[getCurrentExternalInputsStatusParams]

func (s *AVContentService) GetCurrentExternalInputsStatus() (*GetCurrentExternalInputsStatusResult, *http.Response, error) {
	body := getCurrentExternalInputsStatusPayload{
		Method:  "getCurrentExternalInputsStatus",
		ID:      1,
		Params:  getCurrentExternalInputsStatusParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, avContentPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetCurrentExternalInputsStatusResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

type SetPlayContentResult = Result[[0]struct{}]

type setPlayContentParams [1]struct {
	URI string `json:"uri"`
}
type setPlayContentPayload Payload[setPlayContentParams]

func (s *AVContentService) SetPlayContent(uri string) (*SetPlayContentResult, *http.Response, error) {
	body := setPlayContentPayload{
		Method:  "setPlayContent",
		ID:      1,
		Params:  setPlayContentParams{{URI: uri}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, avContentPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(SetPlayContentResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}
