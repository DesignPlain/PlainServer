package cloudwatch

import types "DesignSphere_Server/src/resource/aws/types"

type EventConnection struct {
	// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
	AuthParameters types.Cloudwatch_EventConnectionAuthParameters `json:"authParameters,omitempty" yaml:"authParameters,omitempty"`

	// Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
	AuthorizationType string `json:"authorizationType,omitempty" yaml:"authorizationType,omitempty"`

	// Enter a description for the connection. Maximum of 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
