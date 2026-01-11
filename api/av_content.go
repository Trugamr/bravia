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

// Scheme represents a content scheme
type Scheme struct {
	Scheme string `json:"scheme"`
}

// GetSchemeListResult is the response from the getSchemeList method
type GetSchemeListResult = Result[[1][]Scheme]

type getSchemeListParams [0]struct{}
type getSchemeListPayload Payload[getSchemeListParams]

// GetSchemeList returns the list of available content schemes
func (s *AVContentService) GetSchemeList() (*GetSchemeListResult, *http.Response, error) {
	body := getSchemeListPayload{
		Method:  "getSchemeList",
		ID:      1,
		Params:  getSchemeListParams{},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, avContentPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetSchemeListResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// Source represents a content source
type Source struct {
	Source string `json:"source"`
}

// GetSourceListResult is the response from the getSourceList method
type GetSourceListResult = Result[[1][]Source]

type getSourceListParams [1]struct {
	Scheme string `json:"scheme"`
}
type getSourceListPayload Payload[getSourceListParams]

// GetSourceList returns the list of sources for a given scheme
func (s *AVContentService) GetSourceList(scheme string) (*GetSourceListResult, *http.Response, error) {
	body := getSourceListPayload{
		Method:  "getSourceList",
		ID:      1,
		Params:  getSourceListParams{{Scheme: scheme}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, avContentPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetSourceListResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// ContentCount represents content count information
type ContentCount struct {
	Count int `json:"count"`
}

// GetContentCountResult is the response from the getContentCount method
type GetContentCountResult = Result[[1]ContentCount]

type getContentCountParams [1]struct {
	Source string  `json:"source"`
	Type   *string `json:"type,omitempty"`
}
type getContentCountPayload Payload[getContentCountParams]

// GetContentCount returns the count of content items for a given source
func (s *AVContentService) GetContentCount(source string, contentType *string) (*GetContentCountResult, *http.Response, error) {
	body := getContentCountPayload{
		Method:  "getContentCount",
		ID:      1,
		Params:  getContentCountParams{{Source: source, Type: contentType}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, avContentPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetContentCountResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}

// ContentItem represents a content item
type ContentItem struct {
	URI   string `json:"uri"`
	Title string `json:"title"`
	Index int    `json:"index"`
	// Extended fields for detailed content
	DispNum          *string `json:"dispNum,omitempty"`
	OriginalDispNum  *string `json:"originalDispNum,omitempty"`
	TripletStr       *string `json:"tripletStr,omitempty"`
	ProgramNum       *int    `json:"programNum,omitempty"`
	ProgramMediaType *string `json:"programMediaType,omitempty"`
	DirectRemoteNum  *int    `json:"directRemoteNum,omitempty"`
	StartDateTime    *string `json:"startDateTime,omitempty"`
	DurationSec      *int    `json:"durationSec,omitempty"`
	ChannelName      *string `json:"channelName,omitempty"`
	FileSizeByte     *int    `json:"fileSizeByte,omitempty"`
	IsProtected      *bool   `json:"isProtected,omitempty"`
	IsAlreadyPlayed  *bool   `json:"isAlreadyPlayed,omitempty"`
}

// GetContentListResult is the response from the getContentList method
type GetContentListResult = Result[[1][]ContentItem]

type getContentListParams [1]struct {
	Source string  `json:"source"`
	StIdx  *int    `json:"stIdx,omitempty"`
	Cnt    *int    `json:"cnt,omitempty"`
	Type   *string `json:"type,omitempty"`
}
type getContentListPayload Payload[getContentListParams]

// GetContentList returns a list of content items for a given source
func (s *AVContentService) GetContentList(source string, startIndex, count *int, contentType *string) (*GetContentListResult, *http.Response, error) {
	body := getContentListPayload{
		Method:  "getContentList",
		ID:      1,
		Params:  getContentListParams{{Source: source, StIdx: startIndex, Cnt: count, Type: contentType}},
		Version: "1.0",
	}

	req, err := s.client.NewRequest(http.MethodPost, avContentPath, body)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetContentListResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	if result.HasError() {
		return result, resp, errors.New(result.ErrorMessage())
	}

	return result, resp, nil
}
