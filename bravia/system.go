package bravia

import (
	"net/http"
)

const (
	systemPath = "/sony/system"
)

// SystemService handles requests related to system, such as power status
type SystemService service

// SetPowerStatusResult is the response from the setPowerStatus method
type SetPowerStatusResult Result[[]struct{}]

type SetPowerStatusParam []struct {
	Status bool `json:"status"`
}
type SetPowerStatusPayload Payload[SetPowerStatusParam]

func (s *SystemService) SetPowerStatus(status bool) (*SetPowerStatusResult, *http.Response, error) {
	body := SetPowerStatusPayload{
		Method:  "setPowerStatus",
		ID:      1,
		Params:  SetPowerStatusParam{{Status: status}},
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

	return result, resp, nil
}

// GetPowerStatusResult is the response from the getPowerStatus method
type GetPowerStatusResult Result[[]struct {
	Status string `json:"status"`
}]

// GetPowerStatus returns the power status of the TV
func (s *SystemService) GetPowerStatus() (*GetPowerStatusResult, *http.Response, error) {
	body := Payload[[]struct{}]{
		Method:  "getPowerStatus",
		ID:      1,
		Params:  []struct{}{},
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

	return result, resp, nil
}
