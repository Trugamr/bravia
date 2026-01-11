package api

import (
	"errors"
	"net/http"
)

const (
	audioPath = "/sony/audio"
)

// AudioService handles requests related to audio, such as setting the volume and muting
type AudioService service

type SetAudioVolumeResult = Result[[1]int]

type setAudioVolumeParams [1]struct {
	// Volume specifies the volume level to set.
	// The following formats are accepted:
	//   "N"   - Sets the volume to level N, where N is a numeric string (e.g., "25").
	//   "+N"  - Increases the volume by N, where N is a numeric string (e.g., "+14").
	//   "-N"  - Decreases the volume by N, where N is a numeric string (e.g., "-10").
	Volume string `json:"volume"`

	Target string `json:"target"`
}

type setAudioVolumePayload Payload[setAudioVolumeParams]

// SetAudioVolume sets the volume of the TV
func (s *AudioService) SetAudioVolume(volume, target string) (*SetAudioVolumeResult, *http.Response, error) {
	body := setAudioVolumePayload{
		Method:  "setAudioVolume",
		ID:      1,
		Params:  setAudioVolumeParams{{Volume: volume, Target: target}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, audioPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(SetAudioVolumeResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// VolumeInfo represents volume information for a target
type VolumeInfo struct {
	Target    string `json:"target"`
	Volume    int    `json:"volume"`
	Mute      bool   `json:"mute"`
	MaxVolume int    `json:"maxVolume"`
	MinVolume int    `json:"minVolume"`
}

// GetVolumeInformationResult is the response from the getVolumeInformation method
type GetVolumeInformationResult = Result[[1][]VolumeInfo]

type getVolumeInformationParams [0]struct{}
type getVolumeInformationPayload Payload[getVolumeInformationParams]

// GetVolumeInformation returns volume information for all targets
func (s *AudioService) GetVolumeInformation() (*GetVolumeInformationResult, *http.Response, error) {
	body := getVolumeInformationPayload{
		Method:  "getVolumeInformation",
		ID:      1,
		Params:  getVolumeInformationParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, audioPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetVolumeInformationResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// SetAudioMuteResult is the response from the setAudioMute method
type SetAudioMuteResult = Result[[1]bool]

type setAudioMuteParams [1]struct {
	Status bool `json:"status"`
}
type setAudioMutePayload Payload[setAudioMuteParams]

// SetAudioMute sets the audio mute status
func (s *AudioService) SetAudioMute(status bool) (*SetAudioMuteResult, *http.Response, error) {
	body := setAudioMutePayload{
		Method:  "setAudioMute",
		ID:      1,
		Params:  setAudioMuteParams{{Status: status}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, audioPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(SetAudioMuteResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}
