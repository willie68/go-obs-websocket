package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// StartStopRecordingRequest : Toggle recording on or off.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstoprecording
type StartStopRecordingRequest struct {
	_request `json:",squash"`
	response chan StartStopRecordingResponse
}

// NewStartStopRecordingRequest returns a new StartStopRecordingRequest.
func NewStartStopRecordingRequest() StartStopRecordingRequest {
	return StartStopRecordingRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "StartStopRecording",
			err:   make(chan error, 1),
		},
		make(chan StartStopRecordingResponse, 1),
	}
}

// Send sends the request.
func (r *StartStopRecordingRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StartStopRecordingResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StartStopRecordingRequest) Receive() (StartStopRecordingResponse, error) {
	if !r.sent {
		return StartStopRecordingResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStopRecordingResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartStopRecordingResponse{}, err
		case <-time.After(receiveTimeout):
			return StartStopRecordingResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartStopRecordingRequest) SendReceive(c Client) (StartStopRecordingResponse, error) {
	if err := r.Send(c); err != nil {
		return StartStopRecordingResponse{}, err
	}
	return r.Receive()
}

// StartStopRecordingResponse : Response for StartStopRecordingRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startstoprecording
type StartStopRecordingResponse struct {
	_response `json:",squash"`
}

// StartRecordingRequest : Start recording.
// Will return an `error` if recording is already active.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startrecording
type StartRecordingRequest struct {
	_request `json:",squash"`
	response chan StartRecordingResponse
}

// NewStartRecordingRequest returns a new StartRecordingRequest.
func NewStartRecordingRequest() StartRecordingRequest {
	return StartRecordingRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "StartRecording",
			err:   make(chan error, 1),
		},
		make(chan StartRecordingResponse, 1),
	}
}

// Send sends the request.
func (r *StartRecordingRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StartRecordingResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StartRecordingRequest) Receive() (StartRecordingResponse, error) {
	if !r.sent {
		return StartRecordingResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartRecordingResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StartRecordingResponse{}, err
		case <-time.After(receiveTimeout):
			return StartRecordingResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StartRecordingRequest) SendReceive(c Client) (StartRecordingResponse, error) {
	if err := r.Send(c); err != nil {
		return StartRecordingResponse{}, err
	}
	return r.Receive()
}

// StartRecordingResponse : Response for StartRecordingRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#startrecording
type StartRecordingResponse struct {
	_response `json:",squash"`
}

// StopRecordingRequest : Stop recording.
// Will return an `error` if recording is not active.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stoprecording
type StopRecordingRequest struct {
	_request `json:",squash"`
	response chan StopRecordingResponse
}

// NewStopRecordingRequest returns a new StopRecordingRequest.
func NewStopRecordingRequest() StopRecordingRequest {
	return StopRecordingRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "StopRecording",
			err:   make(chan error, 1),
		},
		make(chan StopRecordingResponse, 1),
	}
}

// Send sends the request.
func (r *StopRecordingRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp StopRecordingResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r StopRecordingRequest) Receive() (StopRecordingResponse, error) {
	if !r.sent {
		return StopRecordingResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopRecordingResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopRecordingResponse{}, err
		case <-time.After(receiveTimeout):
			return StopRecordingResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StopRecordingRequest) SendReceive(c Client) (StopRecordingResponse, error) {
	if err := r.Send(c); err != nil {
		return StopRecordingResponse{}, err
	}
	return r.Receive()
}

// StopRecordingResponse : Response for StopRecordingRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stoprecording
type StopRecordingResponse struct {
	_response `json:",squash"`
}

// SetRecordingFolderRequest : Change the current recording folder.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setrecordingfolder
type SetRecordingFolderRequest struct {
	// Path of the recording folder.
	// Required: Yes.
	RecFolder string `json:"rec-folder"`
	_request  `json:",squash"`
	response  chan SetRecordingFolderResponse
}

// NewSetRecordingFolderRequest returns a new SetRecordingFolderRequest.
func NewSetRecordingFolderRequest(recFolder string) SetRecordingFolderRequest {
	return SetRecordingFolderRequest{
		recFolder,
		_request{
			ID_:   getMessageID(),
			Type_: "SetRecordingFolder",
			err:   make(chan error, 1),
		},
		make(chan SetRecordingFolderResponse, 1),
	}
}

// Send sends the request.
func (r *SetRecordingFolderRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp SetRecordingFolderResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r SetRecordingFolderRequest) Receive() (SetRecordingFolderResponse, error) {
	if !r.sent {
		return SetRecordingFolderResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetRecordingFolderResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetRecordingFolderResponse{}, err
		case <-time.After(receiveTimeout):
			return SetRecordingFolderResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetRecordingFolderRequest) SendReceive(c Client) (SetRecordingFolderResponse, error) {
	if err := r.Send(c); err != nil {
		return SetRecordingFolderResponse{}, err
	}
	return r.Receive()
}

// SetRecordingFolderResponse : Response for SetRecordingFolderRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setrecordingfolder
type SetRecordingFolderResponse struct {
	_response `json:",squash"`
}

// GetRecordingFolderRequest : Get the path of  the current recording folder.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getrecordingfolder
type GetRecordingFolderRequest struct {
	_request `json:",squash"`
	response chan GetRecordingFolderResponse
}

// NewGetRecordingFolderRequest returns a new GetRecordingFolderRequest.
func NewGetRecordingFolderRequest() GetRecordingFolderRequest {
	return GetRecordingFolderRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetRecordingFolder",
			err:   make(chan error, 1),
		},
		make(chan GetRecordingFolderResponse, 1),
	}
}

// Send sends the request.
func (r *GetRecordingFolderRequest) Send(c Client) error {
	if r.sent {
		return ErrAlreadySent
	}
	future, err := c.sendRequest(r)
	if err != nil {
		return err
	}
	r.sent = true
	go func() {
		m := <-future
		var resp GetRecordingFolderResponse
		if err = mapToStruct(m, &resp); err != nil {
			r.err <- err
		} else if resp.Status() != StatusOK {
			r.err <- errors.New(resp.Error())
		} else {
			r.response <- resp
		}
	}()
	return nil
}

// Receive waits for the response.
func (r GetRecordingFolderRequest) Receive() (GetRecordingFolderResponse, error) {
	if !r.sent {
		return GetRecordingFolderResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetRecordingFolderResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetRecordingFolderResponse{}, err
		case <-time.After(receiveTimeout):
			return GetRecordingFolderResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetRecordingFolderRequest) SendReceive(c Client) (GetRecordingFolderResponse, error) {
	if err := r.Send(c); err != nil {
		return GetRecordingFolderResponse{}, err
	}
	return r.Receive()
}

// GetRecordingFolderResponse : Response for GetRecordingFolderRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getrecordingfolder
type GetRecordingFolderResponse struct {
	// Path of the recording folder.
	// Required: Yes.
	RecFolder string `json:"rec-folder"`
	_response `json:",squash"`
}
