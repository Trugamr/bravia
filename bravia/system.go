package bravia

import (
	"encoding/json"
	"net/http"
)

const (
	systemPath = "/sony/system"
)

// SystemService handles requests related to system, such as power status
type SystemService service

// PowerStatusResponse is the response from the getPowerStatus method
type PowerStatusResponse struct {
	Result []struct {
		Status string `json:"status"`
	} `json:"result"`
	ID int `json:"id"`
}

// GetPowerStatus returns the power status of the TV
func (s *SystemService) GetPowerStatus() (string, error) {
	body := map[string]interface{}{
		"method":  "getPowerStatus",
		"id":      1,
		"params":  []interface{}{},
		"version": "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response PowerStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Result[0].Status, nil
}
