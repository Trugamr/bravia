package bravia

import (
	"net/http"
)

const (
	systemPath = "/sony/system"
)

// SystemService handles requests related to system, such as power status
type SystemService service

// SetPowerStatusResponse is the response from the setPowerStatus method
type SetPowerStatusResponse struct {
	Result []struct{} `json:"result"`
	ID     int        `json:"id"`
}

func (s *SystemService) SetPowerStatus(status bool) (*SetPowerStatusResponse, *http.Response, error) {
	body := map[string]interface{}{
		"method":  "setPowerStatus",
		"id":      1,
		"params":  []map[string]bool{{"status": status}},
		"version": "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(SetPowerStatusResponse)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetPowerStatusResponse is the response from the getPowerStatus method
type GetPowerStatusResponse struct {
	Result []struct {
		Status string `json:"status"`
	} `json:"result"`
	ID int `json:"id"`
}

// GetPowerStatus returns the power status of the TV
func (s *SystemService) GetPowerStatus() (*GetPowerStatusResponse, *http.Response, error) {
	body := map[string]interface{}{
		"method":  "getPowerStatus",
		"id":      1,
		"params":  []interface{}{},
		"version": "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetPowerStatusResponse)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
