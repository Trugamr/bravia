package api

import (
	"errors"
	"net/http"
)

const (
	systemPath = "/sony/system"
)

// SystemService handles requests related to system, such as power status
type SystemService service

// SetPowerStatusResult is the response from the setPowerStatus method
type SetPowerStatusResult = Result[[0]struct{}]

type setPowerStatusParams [1]struct {
	Status bool `json:"status"`
}
type setPowerStatusPayload Payload[setPowerStatusParams]

func (s *SystemService) SetPowerStatus(status bool) (*SetPowerStatusResult, *http.Response, error) {
	body := setPowerStatusPayload{
		Method:  "setPowerStatus",
		ID:      1,
		Params:  setPowerStatusParams{{Status: status}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(SetPowerStatusResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// GetPowerStatusResult is the response from the getPowerStatus method
type GetPowerStatusResult = Result[[1]struct {
	Status string `json:"status"`
}]

type getPowerStatusParams [0]struct{}
type getPowerStatusPayload Payload[getPowerStatusParams]

// GetPowerStatus returns the power status of the TV
func (s *SystemService) GetPowerStatus() (*GetPowerStatusResult, *http.Response, error) {
	body := getPowerStatusPayload{
		Method:  "getPowerStatus",
		ID:      1,
		Params:  getPowerStatusParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetPowerStatusResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}
