package appfabric

import types "libds/aws/types"

type AppAuthorizationConnection struct {
	//
	Timeouts types.Appfabric_AppAuthorizationConnectionTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app authorization to use for the request.
	AppAuthorizationArn string `json:"appAuthorizationArn,omitempty" yaml:"appAuthorizationArn,omitempty"`

	// The Amazon Resource Name (ARN) of the app bundle to use for the request.
	AppBundleArn string `json:"appBundleArn,omitempty" yaml:"appBundleArn,omitempty"`

	// Contains OAuth2 authorization information.This is required if the app authorization for the request is configured with an OAuth2 (oauth2) authorization type.
	AuthRequest types.Appfabric_AppAuthorizationConnectionAuthRequest `json:"authRequest,omitempty" yaml:"authRequest,omitempty"`
}
