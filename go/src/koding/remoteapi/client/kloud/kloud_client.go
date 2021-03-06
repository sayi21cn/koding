package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new kloud API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kloud API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIKloudAddAdmin post remote API kloud add admin API
*/
func (a *Client) PostRemoteAPIKloudAddAdmin(params *PostRemoteAPIKloudAddAdminParams) (*PostRemoteAPIKloudAddAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudAddAdminParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudAddAdmin",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.addAdmin",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudAddAdminReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudAddAdminOK), nil

}

/*
PostRemoteAPIKloudBootstrap post remote API kloud bootstrap API
*/
func (a *Client) PostRemoteAPIKloudBootstrap(params *PostRemoteAPIKloudBootstrapParams) (*PostRemoteAPIKloudBootstrapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudBootstrapParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudBootstrap",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.bootstrap",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudBootstrapReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudBootstrapOK), nil

}

/*
PostRemoteAPIKloudBuild post remote API kloud build API
*/
func (a *Client) PostRemoteAPIKloudBuild(params *PostRemoteAPIKloudBuildParams) (*PostRemoteAPIKloudBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudBuild",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.build",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudBuildReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudBuildOK), nil

}

/*
PostRemoteAPIKloudBuildStack post remote API kloud build stack API
*/
func (a *Client) PostRemoteAPIKloudBuildStack(params *PostRemoteAPIKloudBuildStackParams) (*PostRemoteAPIKloudBuildStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudBuildStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudBuildStack",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.buildStack",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudBuildStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudBuildStackOK), nil

}

/*
PostRemoteAPIKloudCheckCredential post remote API kloud check credential API
*/
func (a *Client) PostRemoteAPIKloudCheckCredential(params *PostRemoteAPIKloudCheckCredentialParams) (*PostRemoteAPIKloudCheckCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudCheckCredentialParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudCheckCredential",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.checkCredential",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudCheckCredentialReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudCheckCredentialOK), nil

}

/*
PostRemoteAPIKloudCheckTemplate post remote API kloud check template API
*/
func (a *Client) PostRemoteAPIKloudCheckTemplate(params *PostRemoteAPIKloudCheckTemplateParams) (*PostRemoteAPIKloudCheckTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudCheckTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudCheckTemplate",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.checkTemplate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudCheckTemplateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudCheckTemplateOK), nil

}

/*
PostRemoteAPIKloudDestroy post remote API kloud destroy API
*/
func (a *Client) PostRemoteAPIKloudDestroy(params *PostRemoteAPIKloudDestroyParams) (*PostRemoteAPIKloudDestroyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudDestroyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudDestroy",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.destroy",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudDestroyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudDestroyOK), nil

}

/*
PostRemoteAPIKloudDestroyStack post remote API kloud destroy stack API
*/
func (a *Client) PostRemoteAPIKloudDestroyStack(params *PostRemoteAPIKloudDestroyStackParams) (*PostRemoteAPIKloudDestroyStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudDestroyStackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudDestroyStack",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.destroyStack",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudDestroyStackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudDestroyStackOK), nil

}

/*
PostRemoteAPIKloudEvent post remote API kloud event API
*/
func (a *Client) PostRemoteAPIKloudEvent(params *PostRemoteAPIKloudEventParams) (*PostRemoteAPIKloudEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudEvent",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.event",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudEventReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudEventOK), nil

}

/*
PostRemoteAPIKloudInfo post remote API kloud info API
*/
func (a *Client) PostRemoteAPIKloudInfo(params *PostRemoteAPIKloudInfoParams) (*PostRemoteAPIKloudInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudInfo",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.info",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudInfoOK), nil

}

/*
PostRemoteAPIKloudMigrate post remote API kloud migrate API
*/
func (a *Client) PostRemoteAPIKloudMigrate(params *PostRemoteAPIKloudMigrateParams) (*PostRemoteAPIKloudMigrateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudMigrateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudMigrate",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.migrate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudMigrateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudMigrateOK), nil

}

/*
PostRemoteAPIKloudPing post remote API kloud ping API
*/
func (a *Client) PostRemoteAPIKloudPing(params *PostRemoteAPIKloudPingParams) (*PostRemoteAPIKloudPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudPing",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.ping",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudPingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudPingOK), nil

}

/*
PostRemoteAPIKloudRemoveAdmin post remote API kloud remove admin API
*/
func (a *Client) PostRemoteAPIKloudRemoveAdmin(params *PostRemoteAPIKloudRemoveAdminParams) (*PostRemoteAPIKloudRemoveAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudRemoveAdminParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudRemoveAdmin",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.removeAdmin",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudRemoveAdminReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudRemoveAdminOK), nil

}

/*
PostRemoteAPIKloudRestart post remote API kloud restart API
*/
func (a *Client) PostRemoteAPIKloudRestart(params *PostRemoteAPIKloudRestartParams) (*PostRemoteAPIKloudRestartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudRestartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudRestart",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.restart",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudRestartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudRestartOK), nil

}

/*
PostRemoteAPIKloudStart post remote API kloud start API
*/
func (a *Client) PostRemoteAPIKloudStart(params *PostRemoteAPIKloudStartParams) (*PostRemoteAPIKloudStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudStartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudStart",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudStartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudStartOK), nil

}

/*
PostRemoteAPIKloudStop post remote API kloud stop API
*/
func (a *Client) PostRemoteAPIKloudStop(params *PostRemoteAPIKloudStopParams) (*PostRemoteAPIKloudStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIKloudStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIKloudStop",
		Method:             "POST",
		PathPattern:        "/remote.api/Kloud.stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIKloudStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIKloudStopOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
