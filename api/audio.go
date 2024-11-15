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
