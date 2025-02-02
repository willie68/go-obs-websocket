package obsws

import (
	"errors"
	"time"
)

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// PlayPauseMediaRequest : Pause or play a media source
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#playpausemedia
type PlayPauseMediaRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// Whether to pause or play the source.
	// `false` for play, `true` for pause.
	// Required: Yes.
	PlayPause bool `json:"playPause"`
	_request  `json:",squash"`
	response  chan PlayPauseMediaResponse
}

// NewPlayPauseMediaRequest returns a new PlayPauseMediaRequest.
func NewPlayPauseMediaRequest(
	sourceName string,
	playPause bool,
) PlayPauseMediaRequest {
	return PlayPauseMediaRequest{
		sourceName,
		playPause,
		_request{
			ID_:   getMessageID(),
			Type_: "PlayPauseMedia",
			err:   make(chan error, 1),
		},
		make(chan PlayPauseMediaResponse, 1),
	}
}

// Send sends the request.
func (r *PlayPauseMediaRequest) Send(c Client) error {
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
		var resp PlayPauseMediaResponse
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
func (r PlayPauseMediaRequest) Receive() (PlayPauseMediaResponse, error) {
	if !r.sent {
		return PlayPauseMediaResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return PlayPauseMediaResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return PlayPauseMediaResponse{}, err
		case <-time.After(receiveTimeout):
			return PlayPauseMediaResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r PlayPauseMediaRequest) SendReceive(c Client) (PlayPauseMediaResponse, error) {
	if err := r.Send(c); err != nil {
		return PlayPauseMediaResponse{}, err
	}
	return r.Receive()
}

// PlayPauseMediaResponse : Response for PlayPauseMediaRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#playpausemedia
type PlayPauseMediaResponse struct {
	_response `json:",squash"`
}

// RestartMediaRequest : Restart a media source
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#restartmedia
type RestartMediaRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan RestartMediaResponse
}

// NewRestartMediaRequest returns a new RestartMediaRequest.
func NewRestartMediaRequest(sourceName string) RestartMediaRequest {
	return RestartMediaRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "RestartMedia",
			err:   make(chan error, 1),
		},
		make(chan RestartMediaResponse, 1),
	}
}

// Send sends the request.
func (r *RestartMediaRequest) Send(c Client) error {
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
		var resp RestartMediaResponse
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
func (r RestartMediaRequest) Receive() (RestartMediaResponse, error) {
	if !r.sent {
		return RestartMediaResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return RestartMediaResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return RestartMediaResponse{}, err
		case <-time.After(receiveTimeout):
			return RestartMediaResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r RestartMediaRequest) SendReceive(c Client) (RestartMediaResponse, error) {
	if err := r.Send(c); err != nil {
		return RestartMediaResponse{}, err
	}
	return r.Receive()
}

// RestartMediaResponse : Response for RestartMediaRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#restartmedia
type RestartMediaResponse struct {
	_response `json:",squash"`
}

// StopMediaRequest : Stop a media source
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopmedia
type StopMediaRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan StopMediaResponse
}

// NewStopMediaRequest returns a new StopMediaRequest.
func NewStopMediaRequest(sourceName string) StopMediaRequest {
	return StopMediaRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "StopMedia",
			err:   make(chan error, 1),
		},
		make(chan StopMediaResponse, 1),
	}
}

// Send sends the request.
func (r *StopMediaRequest) Send(c Client) error {
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
		var resp StopMediaResponse
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
func (r StopMediaRequest) Receive() (StopMediaResponse, error) {
	if !r.sent {
		return StopMediaResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopMediaResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return StopMediaResponse{}, err
		case <-time.After(receiveTimeout):
			return StopMediaResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r StopMediaRequest) SendReceive(c Client) (StopMediaResponse, error) {
	if err := r.Send(c); err != nil {
		return StopMediaResponse{}, err
	}
	return r.Receive()
}

// StopMediaResponse : Response for StopMediaRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#stopmedia
type StopMediaResponse struct {
	_response `json:",squash"`
}

// NextMediaRequest : Skip to the next media item in the playlist
// Supports only vlc media source (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#nextmedia
type NextMediaRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan NextMediaResponse
}

// NewNextMediaRequest returns a new NextMediaRequest.
func NewNextMediaRequest(sourceName string) NextMediaRequest {
	return NextMediaRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "NextMedia",
			err:   make(chan error, 1),
		},
		make(chan NextMediaResponse, 1),
	}
}

// Send sends the request.
func (r *NextMediaRequest) Send(c Client) error {
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
		var resp NextMediaResponse
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
func (r NextMediaRequest) Receive() (NextMediaResponse, error) {
	if !r.sent {
		return NextMediaResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return NextMediaResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return NextMediaResponse{}, err
		case <-time.After(receiveTimeout):
			return NextMediaResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r NextMediaRequest) SendReceive(c Client) (NextMediaResponse, error) {
	if err := r.Send(c); err != nil {
		return NextMediaResponse{}, err
	}
	return r.Receive()
}

// NextMediaResponse : Response for NextMediaRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#nextmedia
type NextMediaResponse struct {
	_response `json:",squash"`
}

// PreviousMediaRequest : Go to the previous media item in the playlist
// Supports only vlc media source (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#previousmedia
type PreviousMediaRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan PreviousMediaResponse
}

// NewPreviousMediaRequest returns a new PreviousMediaRequest.
func NewPreviousMediaRequest(sourceName string) PreviousMediaRequest {
	return PreviousMediaRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "PreviousMedia",
			err:   make(chan error, 1),
		},
		make(chan PreviousMediaResponse, 1),
	}
}

// Send sends the request.
func (r *PreviousMediaRequest) Send(c Client) error {
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
		var resp PreviousMediaResponse
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
func (r PreviousMediaRequest) Receive() (PreviousMediaResponse, error) {
	if !r.sent {
		return PreviousMediaResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return PreviousMediaResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return PreviousMediaResponse{}, err
		case <-time.After(receiveTimeout):
			return PreviousMediaResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r PreviousMediaRequest) SendReceive(c Client) (PreviousMediaResponse, error) {
	if err := r.Send(c); err != nil {
		return PreviousMediaResponse{}, err
	}
	return r.Receive()
}

// PreviousMediaResponse : Response for PreviousMediaRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#previousmedia
type PreviousMediaResponse struct {
	_response `json:",squash"`
}

// GetMediaDurationRequest : Get the length of media in milliseconds
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
// Note: For some reason, for the first 5 or so seconds that the media is playing, the total duration can be off by upwards of 50ms.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getmediaduration
type GetMediaDurationRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan GetMediaDurationResponse
}

// NewGetMediaDurationRequest returns a new GetMediaDurationRequest.
func NewGetMediaDurationRequest(sourceName string) GetMediaDurationRequest {
	return GetMediaDurationRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "GetMediaDuration",
			err:   make(chan error, 1),
		},
		make(chan GetMediaDurationResponse, 1),
	}
}

// Send sends the request.
func (r *GetMediaDurationRequest) Send(c Client) error {
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
		var resp GetMediaDurationResponse
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
func (r GetMediaDurationRequest) Receive() (GetMediaDurationResponse, error) {
	if !r.sent {
		return GetMediaDurationResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetMediaDurationResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetMediaDurationResponse{}, err
		case <-time.After(receiveTimeout):
			return GetMediaDurationResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetMediaDurationRequest) SendReceive(c Client) (GetMediaDurationResponse, error) {
	if err := r.Send(c); err != nil {
		return GetMediaDurationResponse{}, err
	}
	return r.Receive()
}

// GetMediaDurationResponse : Response for GetMediaDurationRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getmediaduration
type GetMediaDurationResponse struct {
	// The total length of media in milliseconds..
	// Required: Yes.
	MediaDuration int `json:"mediaDuration"`
	_response     `json:",squash"`
}

// GetMediaTimeRequest : Get the current timestamp of media in milliseconds
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getmediatime
type GetMediaTimeRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan GetMediaTimeResponse
}

// NewGetMediaTimeRequest returns a new GetMediaTimeRequest.
func NewGetMediaTimeRequest(sourceName string) GetMediaTimeRequest {
	return GetMediaTimeRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "GetMediaTime",
			err:   make(chan error, 1),
		},
		make(chan GetMediaTimeResponse, 1),
	}
}

// Send sends the request.
func (r *GetMediaTimeRequest) Send(c Client) error {
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
		var resp GetMediaTimeResponse
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
func (r GetMediaTimeRequest) Receive() (GetMediaTimeResponse, error) {
	if !r.sent {
		return GetMediaTimeResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetMediaTimeResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetMediaTimeResponse{}, err
		case <-time.After(receiveTimeout):
			return GetMediaTimeResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetMediaTimeRequest) SendReceive(c Client) (GetMediaTimeResponse, error) {
	if err := r.Send(c); err != nil {
		return GetMediaTimeResponse{}, err
	}
	return r.Receive()
}

// GetMediaTimeResponse : Response for GetMediaTimeRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getmediatime
type GetMediaTimeResponse struct {
	// The time in milliseconds since the start of the media.
	// Required: Yes.
	Timestamp int `json:"timestamp"`
	_response `json:",squash"`
}

// SetMediaTimeRequest : Set the timestamp of a media source
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setmediatime
type SetMediaTimeRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// Milliseconds to set the timestamp to.
	// Required: Yes.
	Timestamp int `json:"timestamp"`
	_request  `json:",squash"`
	response  chan SetMediaTimeResponse
}

// NewSetMediaTimeRequest returns a new SetMediaTimeRequest.
func NewSetMediaTimeRequest(
	sourceName string,
	timestamp int,
) SetMediaTimeRequest {
	return SetMediaTimeRequest{
		sourceName,
		timestamp,
		_request{
			ID_:   getMessageID(),
			Type_: "SetMediaTime",
			err:   make(chan error, 1),
		},
		make(chan SetMediaTimeResponse, 1),
	}
}

// Send sends the request.
func (r *SetMediaTimeRequest) Send(c Client) error {
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
		var resp SetMediaTimeResponse
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
func (r SetMediaTimeRequest) Receive() (SetMediaTimeResponse, error) {
	if !r.sent {
		return SetMediaTimeResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetMediaTimeResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return SetMediaTimeResponse{}, err
		case <-time.After(receiveTimeout):
			return SetMediaTimeResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r SetMediaTimeRequest) SendReceive(c Client) (SetMediaTimeResponse, error) {
	if err := r.Send(c); err != nil {
		return SetMediaTimeResponse{}, err
	}
	return r.Receive()
}

// SetMediaTimeResponse : Response for SetMediaTimeRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#setmediatime
type SetMediaTimeResponse struct {
	_response `json:",squash"`
}

// ScrubMediaRequest : Scrub media using a supplied offset
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8)
// Note: Due to processing/network delays, this request is not perfect
// The processing rate of this request has also not been tested.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#scrubmedia
type ScrubMediaRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	// Millisecond offset (positive or negative) to offset the current media position.
	// Required: Yes.
	TimeOffset int `json:"timeOffset"`
	_request   `json:",squash"`
	response   chan ScrubMediaResponse
}

// NewScrubMediaRequest returns a new ScrubMediaRequest.
func NewScrubMediaRequest(
	sourceName string,
	timeOffset int,
) ScrubMediaRequest {
	return ScrubMediaRequest{
		sourceName,
		timeOffset,
		_request{
			ID_:   getMessageID(),
			Type_: "ScrubMedia",
			err:   make(chan error, 1),
		},
		make(chan ScrubMediaResponse, 1),
	}
}

// Send sends the request.
func (r *ScrubMediaRequest) Send(c Client) error {
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
		var resp ScrubMediaResponse
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
func (r ScrubMediaRequest) Receive() (ScrubMediaResponse, error) {
	if !r.sent {
		return ScrubMediaResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ScrubMediaResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return ScrubMediaResponse{}, err
		case <-time.After(receiveTimeout):
			return ScrubMediaResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r ScrubMediaRequest) SendReceive(c Client) (ScrubMediaResponse, error) {
	if err := r.Send(c); err != nil {
		return ScrubMediaResponse{}, err
	}
	return r.Receive()
}

// ScrubMediaResponse : Response for ScrubMediaRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#scrubmedia
type ScrubMediaResponse struct {
	_response `json:",squash"`
}

// GetMediaStateRequest : Get the current playing state of a media source
// Supports ffmpeg and vlc media sources (as of OBS v25.0.8).
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getmediastate
type GetMediaStateRequest struct {
	// Source name.
	// Required: Yes.
	SourceName string `json:"sourceName"`
	_request   `json:",squash"`
	response   chan GetMediaStateResponse
}

// NewGetMediaStateRequest returns a new GetMediaStateRequest.
func NewGetMediaStateRequest(sourceName string) GetMediaStateRequest {
	return GetMediaStateRequest{
		sourceName,
		_request{
			ID_:   getMessageID(),
			Type_: "GetMediaState",
			err:   make(chan error, 1),
		},
		make(chan GetMediaStateResponse, 1),
	}
}

// Send sends the request.
func (r *GetMediaStateRequest) Send(c Client) error {
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
		var resp GetMediaStateResponse
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
func (r GetMediaStateRequest) Receive() (GetMediaStateResponse, error) {
	if !r.sent {
		return GetMediaStateResponse{}, ErrNotSent
	}
	if receiveTimeout == 0 {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetMediaStateResponse{}, err
		}
	} else {
		select {
		case resp := <-r.response:
			return resp, nil
		case err := <-r.err:
			return GetMediaStateResponse{}, err
		case <-time.After(receiveTimeout):
			return GetMediaStateResponse{}, ErrReceiveTimeout
		}
	}
}

// SendReceive sends the request then immediately waits for the response.
func (r GetMediaStateRequest) SendReceive(c Client) (GetMediaStateResponse, error) {
	if err := r.Send(c); err != nil {
		return GetMediaStateResponse{}, err
	}
	return r.Receive()
}

// GetMediaStateResponse : Response for GetMediaStateRequest.
//
// Since obs-websocket version: 4.9.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#getmediastate
type GetMediaStateResponse struct {
	// The media state of the provided source.
	// States: `none`, `playing`, `opening`, `buffering`, `paused`, `stopped`, `ended`, `error`, `unknown`.
	// Required: Yes.
	MediaState string `json:"mediaState"`
	_response  `json:",squash"`
}
