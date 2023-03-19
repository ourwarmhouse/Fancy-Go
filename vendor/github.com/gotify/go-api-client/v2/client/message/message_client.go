// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new message API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for message API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateMessage creates a message

__NOTE__: This API ONLY accepts an application token as authentication.
*/
func (a *Client) CreateMessage(params *CreateMessageParams, authInfo runtime.ClientAuthInfoWriter) (*CreateMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMessage",
		Method:             "POST",
		PathPattern:        "/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateMessageOK), nil

}

/*
DeleteAppMessages deletes all messages from a specific application
*/
func (a *Client) DeleteAppMessages(params *DeleteAppMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAppMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAppMessages",
		Method:             "DELETE",
		PathPattern:        "/application/{id}/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppMessagesOK), nil

}

/*
DeleteMessage deletes a message with an id
*/
func (a *Client) DeleteMessage(params *DeleteMessageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMessageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMessage",
		Method:             "DELETE",
		PathPattern:        "/message/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMessageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMessageOK), nil

}

/*
DeleteMessages deletes all messages
*/
func (a *Client) DeleteMessages(params *DeleteMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMessages",
		Method:             "DELETE",
		PathPattern:        "/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMessagesOK), nil

}

/*
GetAppMessages returns all messages from a specific application
*/
func (a *Client) GetAppMessages(params *GetAppMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetAppMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAppMessages",
		Method:             "GET",
		PathPattern:        "/application/{id}/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppMessagesOK), nil

}

/*
GetMessages returns all messages
*/
func (a *Client) GetMessages(params *GetMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*GetMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMessages",
		Method:             "GET",
		PathPattern:        "/message",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMessagesOK), nil

}

/*
StreamMessages websockets return newly created messages
*/
func (a *Client) StreamMessages(params *StreamMessagesParams, authInfo runtime.ClientAuthInfoWriter) (*StreamMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "streamMessages",
		Method:             "GET",
		PathPattern:        "/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StreamMessagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StreamMessagesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
