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

// GetCurrentTimeResult is the response from the getCurrentTime method
type GetCurrentTimeResult = Result[[1]string]

type getCurrentTimeParams [0]struct{}
type getCurrentTimePayload Payload[getCurrentTimeParams]

// GetCurrentTime returns the current time of the TV
func (s *SystemService) GetCurrentTime() (*GetCurrentTimeResult, *http.Response, error) {
	body := getCurrentTimePayload{
		Method:  "getCurrentTime",
		ID:      1,
		Params:  getCurrentTimeParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetCurrentTimeResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// RemoteCommand represents a remote control command
type RemoteCommand struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetRemoteControllerInfoResult is the response from the getRemoteControllerInfo method
type GetRemoteControllerInfoResult = Result[[2]interface{}]

type getRemoteControllerInfoParams [0]struct{}
type getRemoteControllerInfoPayload Payload[getRemoteControllerInfoParams]

// GetRemoteControllerInfo returns information about the remote controller
func (s *SystemService) GetRemoteControllerInfo() (*GetRemoteControllerInfoResult, *http.Response, error) {
	body := getRemoteControllerInfoPayload{
		Method:  "getRemoteControllerInfo",
		ID:      1,
		Params:  getRemoteControllerInfoParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetRemoteControllerInfoResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// InterfaceInformation represents TV interface information
type InterfaceInformation struct {
	ProductCategory  string `json:"productCategory"`
	ProductName      string `json:"productName"`
	ModelName        string `json:"modelName"`
	ServerName       string `json:"serverName"`
	InterfaceVersion string `json:"interfaceVersion"`
}

// GetInterfaceInformationResult is the response from the getInterfaceInformation method
type GetInterfaceInformationResult = Result[[1]InterfaceInformation]

type getInterfaceInformationParams [0]struct{}
type getInterfaceInformationPayload Payload[getInterfaceInformationParams]

// GetInterfaceInformation returns information about the TV interface
func (s *SystemService) GetInterfaceInformation() (*GetInterfaceInformationResult, *http.Response, error) {
	body := getInterfaceInformationPayload{
		Method:  "getInterfaceInformation",
		ID:      1,
		Params:  getInterfaceInformationParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetInterfaceInformationResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// RequestRebootResult is the response from the requestReboot method
type RequestRebootResult = Result[[0]struct{}]

type requestRebootParams [0]struct{}
type requestRebootPayload Payload[requestRebootParams]

// RequestReboot requests a reboot of the TV
func (s *SystemService) RequestReboot() (*RequestRebootResult, *http.Response, error) {
	body := requestRebootPayload{
		Method:  "requestReboot",
		ID:      1,
		Params:  requestRebootParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, systemPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(RequestRebootResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}
