package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SetCurrentSceneRequest : Switch to the specified scene.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setcurrentscene
type SetCurrentSceneRequest struct {
	// Name of the scene to switch to.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	_request  `json:",squash"`
	response  chan SetCurrentSceneResponse
}

// NewSetCurrentSceneRequest returns a new SetCurrentSceneRequest.
func NewSetCurrentSceneRequest(sceneName string) SetCurrentSceneRequest {
	return SetCurrentSceneRequest{
		sceneName,
		_request{
			ID_:   getMessageID(),
			Type_: "SetCurrentScene",
			err:   make(chan error, 1),
		},
		make(chan SetCurrentSceneResponse, 1),
	}
}

// Send sends the request.
func (r *SetCurrentSceneRequest) Send(c Client) error {
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
		var resp SetCurrentSceneResponse
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
func (r SetCurrentSceneRequest) Receive() (SetCurrentSceneResponse, error) {
	if !r.sent {
		return SetCurrentSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetCurrentSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return SetCurrentSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetCurrentSceneRequest) SendReceive(c Client) (SetCurrentSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return SetCurrentSceneResponse{}, err
	}
	return r.Receive()
}

// SetCurrentSceneResponse : Response for SetCurrentSceneRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setcurrentscene
type SetCurrentSceneResponse struct {
	_response `json:",squash"`
}

// GetCurrentSceneRequest : Get the current scene's name and source items.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getcurrentscene
type GetCurrentSceneRequest struct {
	_request `json:",squash"`
	response chan GetCurrentSceneResponse
}

// NewGetCurrentSceneRequest returns a new GetCurrentSceneRequest.
func NewGetCurrentSceneRequest() GetCurrentSceneRequest {
	return GetCurrentSceneRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetCurrentScene",
			err:   make(chan error, 1),
		},
		make(chan GetCurrentSceneResponse, 1),
	}
}

// Send sends the request.
func (r *GetCurrentSceneRequest) Send(c Client) error {
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
		var resp GetCurrentSceneResponse
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
func (r GetCurrentSceneRequest) Receive() (GetCurrentSceneResponse, error) {
	if !r.sent {
		return GetCurrentSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetCurrentSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return GetCurrentSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetCurrentSceneRequest) SendReceive(c Client) (GetCurrentSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return GetCurrentSceneResponse{}, err
	}
	return r.Receive()
}

// GetCurrentSceneResponse : Response for GetCurrentSceneRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getcurrentscene
type GetCurrentSceneResponse struct {
	// Name of the currently active scene.
	// Required: Yes.
	Name string `json:"name"`
	// Ordered list of the current scene's source items.
	// Required: Yes.
	Sources   []map[string]interface{} `json:"sources"`
	_response `json:",squash"`
}

// GetSceneListRequest : Get a list of scenes in the currently active profile.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getscenelist
type GetSceneListRequest struct {
	_request `json:",squash"`
	response chan GetSceneListResponse
}

// NewGetSceneListRequest returns a new GetSceneListRequest.
func NewGetSceneListRequest() GetSceneListRequest {
	return GetSceneListRequest{
		_request{
			ID_:   getMessageID(),
			Type_: "GetSceneList",
			err:   make(chan error, 1),
		},
		make(chan GetSceneListResponse, 1),
	}
}

// Send sends the request.
func (r *GetSceneListRequest) Send(c Client) error {
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
		var resp GetSceneListResponse
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
func (r GetSceneListRequest) Receive() (GetSceneListResponse, error) {
	if !r.sent {
		return GetSceneListResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneListResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneListResponse{}, err
		case <-time.After(receiveTimeout):
			return GetSceneListResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetSceneListRequest) SendReceive(c Client) (GetSceneListResponse, error) {
	if err := r.Send(c); err != nil {
		return GetSceneListResponse{}, err
	}
	return r.Receive()
}

// GetSceneListResponse : Response for GetSceneListRequest.
//
// Since obs-websocket version: 0.3.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getscenelist
type GetSceneListResponse struct {
	// Name of the currently active scene.
	// Required: Yes.
	CurrentScene string `json:"current-scene"`
	// Ordered list of the current profile's scenes (See [GetCurrentScene](#getcurrentscene) for more information).
	// Required: Yes.
	Scenes    []map[string]interface{} `json:"scenes"`
	_response `json:",squash"`
}

// CreateSceneRequest : Create a new scene scene.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#createscene
type CreateSceneRequest struct {
	// Name of the scene to create.
	// Required: Yes.
	SceneName string `json:"sceneName"`
	_request  `json:",squash"`
	response  chan CreateSceneResponse
}

// NewCreateSceneRequest returns a new CreateSceneRequest.
func NewCreateSceneRequest(sceneName string) CreateSceneRequest {
	return CreateSceneRequest{
		sceneName,
		_request{
			ID_:   getMessageID(),
			Type_: "CreateScene",
			err:   make(chan error, 1),
		},
		make(chan CreateSceneResponse, 1),
	}
}

// Send sends the request.
func (r *CreateSceneRequest) Send(c Client) error {
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
		var resp CreateSceneResponse
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
func (r CreateSceneRequest) Receive() (CreateSceneResponse, error) {
	if !r.sent {
		return CreateSceneResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return CreateSceneResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return CreateSceneResponse{}, err
		case <-time.After(receiveTimeout):
			return CreateSceneResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r CreateSceneRequest) SendReceive(c Client) (CreateSceneResponse, error) {
	if err := r.Send(c); err != nil {
		return CreateSceneResponse{}, err
	}
	return r.Receive()
}

// CreateSceneResponse : Response for CreateSceneRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#createscene
type CreateSceneResponse struct {
	_response `json:",squash"`
}

// ReorderSceneItemsRequest : Changes the order of scene items in the requested scene.
//
// Since obs-websocket version: 4.5.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#reordersceneitems
type ReorderSceneItemsRequest struct {
	// Name of the scene to reorder (defaults to current).
	// Required: No.
	Scene string `json:"scene"`
	// Ordered list of objects with name and/or id specified.
	// Id preferred due to uniqueness per scene.
	// Required: Yes.
	Items []map[string]interface{} `json:"items"`
	// Id of a specific scene item.
	// Unique on a scene by scene basis.
	// Required: No.
	ItemsID int `json:"items.*.id"`
	// Name of a scene item.
	// Sufficiently unique if no scene items share sources within the scene.
	// Required: No.
	ItemsName string `json:"items.*.name"`
	_request  `json:",squash"`
	response  chan ReorderSceneItemsResponse
}

// NewReorderSceneItemsRequest returns a new ReorderSceneItemsRequest.
func NewReorderSceneItemsRequest(
	scene string,
	items []map[string]interface{},
	itemsID int,
	itemsName string,
) ReorderSceneItemsRequest {
	return ReorderSceneItemsRequest{
		scene,
		items,
		itemsID,
		itemsName,
		_request{
			ID_:   getMessageID(),
			Type_: "ReorderSceneItems",
			err:   make(chan error, 1),
		},
		make(chan ReorderSceneItemsResponse, 1),
	}
}

// Send sends the request.
func (r *ReorderSceneItemsRequest) Send(c Client) error {
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
		var resp ReorderSceneItemsResponse
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
func (r ReorderSceneItemsRequest) Receive() (ReorderSceneItemsResponse, error) {
	if !r.sent {
		return ReorderSceneItemsResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ReorderSceneItemsResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ReorderSceneItemsResponse{}, err
		case <-time.After(receiveTimeout):
			return ReorderSceneItemsResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ReorderSceneItemsRequest) SendReceive(c Client) (ReorderSceneItemsResponse, error) {
	if err := r.Send(c); err != nil {
		return ReorderSceneItemsResponse{}, err
	}
	return r.Receive()
}

// ReorderSceneItemsResponse : Response for ReorderSceneItemsRequest.
//
// Since obs-websocket version: 4.5.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#reordersceneitems
type ReorderSceneItemsResponse struct {
	_response `json:",squash"`
}

// SetSceneTransitionOverrideRequest : Set a scene to use a specific transition override.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setscenetransitionoverride
type SetSceneTransitionOverrideRequest struct {
	// Name of the scene to switch to.
	// Required: Yes.
	SceneName string `json:"sceneName"`
	// Name of the transition to use.
	// Required: Yes.
	TransitionName string `json:"transitionName"`
	// Duration in milliseconds of the transition if transition is not fixed.
	// Defaults to the current duration specified in the UI if there is no current override and this value is not given.
	// Required: Yes.
	TransitionDuration int `json:"transitionDuration"`
	_request           `json:",squash"`
	response           chan SetSceneTransitionOverrideResponse
}

// NewSetSceneTransitionOverrideRequest returns a new SetSceneTransitionOverrideRequest.
func NewSetSceneTransitionOverrideRequest(
	sceneName string,
	transitionName string,
	transitionDuration int,
) SetSceneTransitionOverrideRequest {
	return SetSceneTransitionOverrideRequest{
		sceneName,
		transitionName,
		transitionDuration,
		_request{
			ID_:   getMessageID(),
			Type_: "SetSceneTransitionOverride",
			err:   make(chan error, 1),
		},
		make(chan SetSceneTransitionOverrideResponse, 1),
	}
}

// Send sends the request.
func (r *SetSceneTransitionOverrideRequest) Send(c Client) error {
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
		var resp SetSceneTransitionOverrideResponse
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
func (r SetSceneTransitionOverrideRequest) Receive() (SetSceneTransitionOverrideResponse, error) {
	if !r.sent {
		return SetSceneTransitionOverrideResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneTransitionOverrideResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetSceneTransitionOverrideResponse{}, err
		case <-time.After(receiveTimeout):
			return SetSceneTransitionOverrideResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetSceneTransitionOverrideRequest) SendReceive(c Client) (SetSceneTransitionOverrideResponse, error) {
	if err := r.Send(c); err != nil {
		return SetSceneTransitionOverrideResponse{}, err
	}
	return r.Receive()
}

// SetSceneTransitionOverrideResponse : Response for SetSceneTransitionOverrideRequest.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setscenetransitionoverride
type SetSceneTransitionOverrideResponse struct {
	_response `json:",squash"`
}

// RemoveSceneTransitionOverrideRequest : Remove any transition override on a scene.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#removescenetransitionoverride
type RemoveSceneTransitionOverrideRequest struct {
	// Name of the scene to switch to.
	// Required: Yes.
	SceneName string `json:"sceneName"`
	_request  `json:",squash"`
	response  chan RemoveSceneTransitionOverrideResponse
}

// NewRemoveSceneTransitionOverrideRequest returns a new RemoveSceneTransitionOverrideRequest.
func NewRemoveSceneTransitionOverrideRequest(sceneName string) RemoveSceneTransitionOverrideRequest {
	return RemoveSceneTransitionOverrideRequest{
		sceneName,
		_request{
			ID_:   getMessageID(),
			Type_: "RemoveSceneTransitionOverride",
			err:   make(chan error, 1),
		},
		make(chan RemoveSceneTransitionOverrideResponse, 1),
	}
}

// Send sends the request.
func (r *RemoveSceneTransitionOverrideRequest) Send(c Client) error {
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
		var resp RemoveSceneTransitionOverrideResponse
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
func (r RemoveSceneTransitionOverrideRequest) Receive() (RemoveSceneTransitionOverrideResponse, error) {
	if !r.sent {
		return RemoveSceneTransitionOverrideResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return RemoveSceneTransitionOverrideResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return RemoveSceneTransitionOverrideResponse{}, err
		case <-time.After(receiveTimeout):
			return RemoveSceneTransitionOverrideResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r RemoveSceneTransitionOverrideRequest) SendReceive(c Client) (RemoveSceneTransitionOverrideResponse, error) {
	if err := r.Send(c); err != nil {
		return RemoveSceneTransitionOverrideResponse{}, err
	}
	return r.Receive()
}

// RemoveSceneTransitionOverrideResponse : Response for RemoveSceneTransitionOverrideRequest.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#removescenetransitionoverride
type RemoveSceneTransitionOverrideResponse struct {
	_response `json:",squash"`
}

// GetSceneTransitionOverrideRequest : Get the current scene transition override.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getscenetransitionoverride
type GetSceneTransitionOverrideRequest struct {
	// Name of the scene to switch to.
	// Required: Yes.
	SceneName string `json:"sceneName"`
	_request  `json:",squash"`
	response  chan GetSceneTransitionOverrideResponse
}

// NewGetSceneTransitionOverrideRequest returns a new GetSceneTransitionOverrideRequest.
func NewGetSceneTransitionOverrideRequest(sceneName string) GetSceneTransitionOverrideRequest {
	return GetSceneTransitionOverrideRequest{
		sceneName,
		_request{
			ID_:   getMessageID(),
			Type_: "GetSceneTransitionOverride",
			err:   make(chan error, 1),
		},
		make(chan GetSceneTransitionOverrideResponse, 1),
	}
}

// Send sends the request.
func (r *GetSceneTransitionOverrideRequest) Send(c Client) error {
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
		var resp GetSceneTransitionOverrideResponse
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
func (r GetSceneTransitionOverrideRequest) Receive() (GetSceneTransitionOverrideResponse, error) {
	if !r.sent {
		return GetSceneTransitionOverrideResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneTransitionOverrideResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetSceneTransitionOverrideResponse{}, err
		case <-time.After(receiveTimeout):
			return GetSceneTransitionOverrideResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetSceneTransitionOverrideRequest) SendReceive(c Client) (GetSceneTransitionOverrideResponse, error) {
	if err := r.Send(c); err != nil {
		return GetSceneTransitionOverrideResponse{}, err
	}
	return r.Receive()
}

// GetSceneTransitionOverrideResponse : Response for GetSceneTransitionOverrideRequest.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getscenetransitionoverride
type GetSceneTransitionOverrideResponse struct {
	// Name of the current overriding transition.
	// Empty string if no override is set.
	// Required: Yes.
	TransitionName string `json:"transitionName"`
	// Transition duration.
	// `-1` if no override is set.
	// Required: Yes.
	TransitionDuration int `json:"transitionDuration"`
	_response          `json:",squash"`
}
