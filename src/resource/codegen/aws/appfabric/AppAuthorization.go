package appfabric

import types "libds/aws/types"

type AppAuthorization struct {
	/*
	   Contains credentials for the application, such as an API key or OAuth2 client ID and secret.
	   Specify credentials that match the authorization type for your request. For example, if the authorization type for your request is OAuth2 (oauth2), then you should provide only the OAuth2 credentials.
	*/
	Credential types.Appfabric_AppAuthorizationCredential `json:"credential,omitempty" yaml:"credential,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Contains information about an application tenant, such as the application display name and identifier.
	Tenants []types.Appfabric_AppAuthorizationTenant `json:"tenants,omitempty" yaml:"tenants,omitempty"`

	//
	Timeouts types.Appfabric_AppAuthorizationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// The name of the application for valid values see https://docs.aws.amazon.com/appfabric/latest/api/API_CreateAppAuthorization.html.
	App string `json:"app,omitempty" yaml:"app,omitempty"`

	// The Amazon Resource Name (ARN) of the app bundle to use for the request.
	AppBundleArn string `json:"appBundleArn,omitempty" yaml:"appBundleArn,omitempty"`

	// The authorization type for the app authorization valid values are oauth2 and apiKey.
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`
}
