package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// GetTransitionListRequest : List of all transitions available in the frontend's dropdown menu.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionlist
type GetTransitionListRequest struct {
	_request `json:",squash"`
	response chan GetTransitionListResponse
}

// NewGetTransitionListRequest returns a new GetTransitionListRequest.
func NewGetTransitionListRequest() GetTransitionListRequest {
	return GetTransitionListRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetTransitionList",
			err:   make(chan error, 1),
		},
		make(chan GetTransitionListResponse, 1),
	}
}

// Send sends the request.
func (r *GetTransitionListRequest) Send(c Client) error {
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
		var resp GetTransitionListResponse
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
func (r GetTransitionListRequest) Receive() (GetTransitionListResponse, error) {
	if !r.sent {
		return GetTransitionListResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionListResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionListResponse{}, err
		case <-time.After(receiveTimeout):
			return GetTransitionListResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetTransitionListRequest) SendReceive(c Client) (GetTransitionListResponse, error) {
	if err := r.Send(c); err != nil {
		return GetTransitionListResponse{}, err
	}
	return r.Receive()
}

// GetTransitionListResponse : Response for GetTransitionListRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionlist
type GetTransitionListResponse struct {
	// Name of the currently active transition.
	// Required: Yes.
	CurrentTransition string `json:"current-transition"`
	// List of transitions.
	// Required: Yes.
	Transitions []map[string]interface{} `json:"transitions"`
	// Name of the transition.
	// Required: Yes.
	TransitionsName string `json:"transitions.*.name"`
	_response       `json:",squash"`
}

// GetCurrentTransitionRequest : Get the name of the currently selected transition in the frontend's dropdown menu.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getcurrenttransition
type GetCurrentTransitionRequest struct {
	_request `json:",squash"`
	response chan GetCurrentTransitionResponse
}

// NewGetCurrentTransitionRequest returns a new GetCurrentTransitionRequest.
func NewGetCurrentTransitionRequest() GetCurrentTransitionRequest {
	return GetCurrentTransitionRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetCurrentTransition",
			err:   make(chan error, 1),
		},
		make(chan GetCurrentTransitionResponse, 1),
	}
}

// Send sends the request.
func (r *GetCurrentTransitionRequest) Send(c Client) error {
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
		var resp GetCurrentTransitionResponse
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
func (r GetCurrentTransitionRequest) Receive() (GetCurrentTransitionResponse, error) {
	if !r.sent {
		return GetCurrentTransitionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentTransitionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentTransitionResponse{}, err
		case <-time.After(receiveTimeout):
			return GetCurrentTransitionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetCurrentTransitionRequest) SendReceive(c Client) (GetCurrentTransitionResponse, error) {
	if err := r.Send(c); err != nil {
		return GetCurrentTransitionResponse{}, err
	}
	return r.Receive()
}

// GetCurrentTransitionResponse : Response for GetCurrentTransitionRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getcurrenttransition
type GetCurrentTransitionResponse struct {
	// Name of the selected transition.
	// Required: Yes.
	Name string `json:"name"`
	// Transition duration (in milliseconds) if supported by the transition.
	// Required: No.
	Duration  int `json:"duration"`
	_response `json:",squash"`
}

// SetCurrentTransitionRequest : Set the active transition.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setcurrenttransition
type SetCurrentTransitionRequest struct {
	// The name of the transition.
	// Required: Yes.
	TransitionName string `json:"transition-name"`
	_request       `json:",squash"`
	response       chan SetCurrentTransitionResponse
}

// NewSetCurrentTransitionRequest returns a new SetCurrentTransitionRequest.
func NewSetCurrentTransitionRequest(transitionName string) SetCurrentTransitionRequest {
	return SetCurrentTransitionRequest{
		transitionName,
		_request{
			ID_:   getMessageID(),
			Type_: "SetCurrentTransition",
			err:   make(chan error, 1),
		},
		make(chan SetCurrentTransitionResponse, 1),
	}
}

// Send sends the request.
func (r *SetCurrentTransitionRequest) Send(c Client) error {
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
		var resp SetCurrentTransitionResponse
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
func (r SetCurrentTransitionRequest) Receive() (SetCurrentTransitionResponse, error) {
	if !r.sent {
		return SetCurrentTransitionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentTransitionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentTransitionResponse{}, err
		case <-time.After(receiveTimeout):
			return SetCurrentTransitionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetCurrentTransitionRequest) SendReceive(c Client) (SetCurrentTransitionResponse, error) {
	if err := r.Send(c); err != nil {
		return SetCurrentTransitionResponse{}, err
	}
	return r.Receive()
}

// SetCurrentTransitionResponse : Response for SetCurrentTransitionRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setcurrenttransition
type SetCurrentTransitionResponse struct {
	_response `json:",squash"`
}

// SetTransitionDurationRequest : Set the duration of the currently selected transition if supported.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#settransitionduration
type SetTransitionDurationRequest struct {
	// Desired duration of the transition (in milliseconds).
	// Required: Yes.
	Duration int `json:"duration"`
	_request `json:",squash"`
	response chan SetTransitionDurationResponse
}

// NewSetTransitionDurationRequest returns a new SetTransitionDurationRequest.
func NewSetTransitionDurationRequest(duration int) SetTransitionDurationRequest {
	return SetTransitionDurationRequest{
		duration,
		_request{
			ID_:   getMessageID(),
			Type_: "SetTransitionDuration",
			err:   make(chan error, 1),
		},
		make(chan SetTransitionDurationResponse, 1),
	}
}

// Send sends the request.
func (r *SetTransitionDurationRequest) Send(c Client) error {
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
		var resp SetTransitionDurationResponse
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
func (r SetTransitionDurationRequest) Receive() (SetTransitionDurationResponse, error) {
	if !r.sent {
		return SetTransitionDurationResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetTransitionDurationResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetTransitionDurationResponse{}, err
		case <-time.After(receiveTimeout):
			return SetTransitionDurationResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetTransitionDurationRequest) SendReceive(c Client) (SetTransitionDurationResponse, error) {
	if err := r.Send(c); err != nil {
		return SetTransitionDurationResponse{}, err
	}
	return r.Receive()
}

// SetTransitionDurationResponse : Response for SetTransitionDurationRequest.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#settransitionduration
type SetTransitionDurationResponse struct {
	_response `json:",squash"`
}

// GetTransitionDurationRequest : Get the duration of the currently selected transition if supported.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionduration
type GetTransitionDurationRequest struct {
	_request `json:",squash"`
	response chan GetTransitionDurationResponse
}

// NewGetTransitionDurationRequest returns a new GetTransitionDurationRequest.
func NewGetTransitionDurationRequest() GetTransitionDurationRequest {
	return GetTransitionDurationRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetTransitionDuration",
			err:   make(chan error, 1),
		},
		make(chan GetTransitionDurationResponse, 1),
	}
}

// Send sends the request.
func (r *GetTransitionDurationRequest) Send(c Client) error {
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
		var resp GetTransitionDurationResponse
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
func (r GetTransitionDurationRequest) Receive() (GetTransitionDurationResponse, error) {
	if !r.sent {
		return GetTransitionDurationResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionDurationResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionDurationResponse{}, err
		case <-time.After(receiveTimeout):
			return GetTransitionDurationResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetTransitionDurationRequest) SendReceive(c Client) (GetTransitionDurationResponse, error) {
	if err := r.Send(c); err != nil {
		return GetTransitionDurationResponse{}, err
	}
	return r.Receive()
}

// GetTransitionDurationResponse : Response for GetTransitionDurationRequest.
//
// Since obs-websocket version: 4.1.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionduration
type GetTransitionDurationResponse struct {
	// Duration of the current transition (in milliseconds).
	// Required: Yes.
	TransitionDuration int `json:"transition-duration"`
	_response          `json:",squash"`
}

// GetTransitionPositionRequest : Get the position of the current transition.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionposition
type GetTransitionPositionRequest struct {
	_request `json:",squash"`
	response chan GetTransitionPositionResponse
}

// NewGetTransitionPositionRequest returns a new GetTransitionPositionRequest.
func NewGetTransitionPositionRequest() GetTransitionPositionRequest {
	return GetTransitionPositionRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetTransitionPosition",
			err:   make(chan error, 1),
		},
		make(chan GetTransitionPositionResponse, 1),
	}
}

// Send sends the request.
func (r *GetTransitionPositionRequest) Send(c Client) error {
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
		var resp GetTransitionPositionResponse
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
func (r GetTransitionPositionRequest) Receive() (GetTransitionPositionResponse, error) {
	if !r.sent {
		return GetTransitionPositionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionPositionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionPositionResponse{}, err
		case <-time.After(receiveTimeout):
			return GetTransitionPositionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetTransitionPositionRequest) SendReceive(c Client) (GetTransitionPositionResponse, error) {
	if err := r.Send(c); err != nil {
		return GetTransitionPositionResponse{}, err
	}
	return r.Receive()
}

// GetTransitionPositionResponse : Response for GetTransitionPositionRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionposition
type GetTransitionPositionResponse struct {
	// current transition position.
	// This value will be between 0.0 and 1.0.
	// Note: Transition returns 1.0 when not active.
	// Required: Yes.
	Position  float64 `json:"position"`
	_response `json:",squash"`
}

// GetTransitionSettingsRequest : Get the current settings of a transition.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionsettings
type GetTransitionSettingsRequest struct {
	// Transition name.
	// Required: Yes.
	TransitionName string `json:"transitionName"`
	_request       `json:",squash"`
	response       chan GetTransitionSettingsResponse
}

// NewGetTransitionSettingsRequest returns a new GetTransitionSettingsRequest.
func NewGetTransitionSettingsRequest(transitionName string) GetTransitionSettingsRequest {
	return GetTransitionSettingsRequest{
		transitionName,
		_request{
			ID_:   getMessageID(),
			Type_: "GetTransitionSettings",
			err:   make(chan error, 1),
		},
		make(chan GetTransitionSettingsResponse, 1),
	}
}

// Send sends the request.
func (r *GetTransitionSettingsRequest) Send(c Client) error {
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
		var resp GetTransitionSettingsResponse
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
func (r GetTransitionSettingsRequest) Receive() (GetTransitionSettingsResponse, error) {
	if !r.sent {
		return GetTransitionSettingsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionSettingsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetTransitionSettingsResponse{}, err
		case <-time.After(receiveTimeout):
			return GetTransitionSettingsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetTransitionSettingsRequest) SendReceive(c Client) (GetTransitionSettingsResponse, error) {
	if err := r.Send(c); err != nil {
		return GetTransitionSettingsResponse{}, err
	}
	return r.Receive()
}

// GetTransitionSettingsResponse : Response for GetTransitionSettingsRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#gettransitionsettings
type GetTransitionSettingsResponse struct {
	// Current transition settings.
	// Required: Yes.
	TransitionSettings map[string]interface{} `json:"transitionSettings"`
	_response          `json:",squash"`
}

// SetTransitionSettingsRequest : Change the current settings of a transition.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#settransitionsettings
type SetTransitionSettingsRequest struct {
	// Transition name.
	// Required: Yes.
	TransitionName string `json:"transitionName"`
	// Transition settings (they can be partial).
	// Required: Yes.
	TransitionSettings map[string]interface{} `json:"transitionSettings"`
	_request           `json:",squash"`
	response           chan SetTransitionSettingsResponse
}

// NewSetTransitionSettingsRequest returns a new SetTransitionSettingsRequest.
func NewSetTransitionSettingsRequest(
	transitionName string,
	transitionSettings map[string]interface{},
) SetTransitionSettingsRequest {
	return SetTransitionSettingsRequest{
		transitionName,
		transitionSettings,
		_request{
			ID_:   getMessageID(),
			Type_: "SetTransitionSettings",
			err:   make(chan error, 1),
		},
		make(chan SetTransitionSettingsResponse, 1),
	}
}

// Send sends the request.
func (r *SetTransitionSettingsRequest) Send(c Client) error {
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
		var resp SetTransitionSettingsResponse
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
func (r SetTransitionSettingsRequest) Receive() (SetTransitionSettingsResponse, error) {
	if !r.sent {
		return SetTransitionSettingsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetTransitionSettingsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetTransitionSettingsResponse{}, err
		case <-time.After(receiveTimeout):
			return SetTransitionSettingsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetTransitionSettingsRequest) SendReceive(c Client) (SetTransitionSettingsResponse, error) {
	if err := r.Send(c); err != nil {
		return SetTransitionSettingsResponse{}, err
	}
	return r.Receive()
}

// SetTransitionSettingsResponse : Response for SetTransitionSettingsRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#settransitionsettings
type SetTransitionSettingsResponse struct {
	// Updated transition settings.
	// Required: Yes.
	TransitionSettings map[string]interface{} `json:"transitionSettings"`
	_response          `json:",squash"`
}

// ReleaseTBarRequest : Release the T-Bar (like a user releasing their mouse button after moving it).
// *YOU MUST CALL THIS if you called `SetTBarPosition` with the `release` parameter set to `false`.*.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#releasetbar
type ReleaseTBarRequest struct {
	_request `json:",squash"`
	response chan ReleaseTBarResponse
}

// NewReleaseTBarRequest returns a new ReleaseTBarRequest.
func NewReleaseTBarRequest() ReleaseTBarRequest {
	return ReleaseTBarRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "ReleaseTBar",
			err:   make(chan error, 1),
		},
		make(chan ReleaseTBarResponse, 1),
	}
}

// Send sends the request.
func (r *ReleaseTBarRequest) Send(c Client) error {
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
		var resp ReleaseTBarResponse
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
func (r ReleaseTBarRequest) Receive() (ReleaseTBarResponse, error) {
	if !r.sent {
		return ReleaseTBarResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ReleaseTBarResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ReleaseTBarResponse{}, err
		case <-time.After(receiveTimeout):
			return ReleaseTBarResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ReleaseTBarRequest) SendReceive(c Client) (ReleaseTBarResponse, error) {
	if err := r.Send(c); err != nil {
		return ReleaseTBarResponse{}, err
	}
	return r.Receive()
}

// ReleaseTBarResponse : Response for ReleaseTBarRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#releasetbar
type ReleaseTBarResponse struct {
	_response `json:",squash"`
}

// SetTBarPositionRequest :
//
// If your code needs to perform multiple successive T-Bar moves (e.g. : in an animation, or in response to a user moving a T-Bar control in your User Interface), set `release` to false and call `ReleaseTBar` later once the animation/interaction is over.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#settbarposition
type SetTBarPositionRequest struct {
	// T-Bar position.
	// This value must be between 0.0 and 1.0.
	// Required: Yes.
	Position float64 `json:"position"`
	// Whether or not the T-Bar gets released automatically after setting its new position (like a user releasing their mouse button after moving the T-Bar).
	// Call `ReleaseTBar` manually if you set `release` to false.
	// Defaults to true.
	// Required: No.
	Release  bool `json:"release"`
	_request `json:",squash"`
	response chan SetTBarPositionResponse
}

// NewSetTBarPositionRequest returns a new SetTBarPositionRequest.
func NewSetTBarPositionRequest(
	position float64,
	release bool,
) SetTBarPositionRequest {
	return SetTBarPositionRequest{
		position,
		release,
		_request{
			ID_:   getMessageID(),
			Type_: "SetTBarPosition",
			err:   make(chan error, 1),
		},
		make(chan SetTBarPositionResponse, 1),
	}
}

// Send sends the request.
func (r *SetTBarPositionRequest) Send(c Client) error {
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
		var resp SetTBarPositionResponse
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
func (r SetTBarPositionRequest) Receive() (SetTBarPositionResponse, error) {
	if !r.sent {
		return SetTBarPositionResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetTBarPositionResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetTBarPositionResponse{}, err
		case <-time.After(receiveTimeout):
			return SetTBarPositionResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetTBarPositionRequest) SendReceive(c Client) (SetTBarPositionResponse, error) {
	if err := r.Send(c); err != nil {
		return SetTBarPositionResponse{}, err
	}
	return r.Receive()
}

// SetTBarPositionResponse : Response for SetTBarPositionRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#settbarposition
type SetTBarPositionResponse struct {
	_response `json:",squash"`
}
