package cognito

import types "DesignSphere_Server/src/resource/aws/types"

type ResourceServer struct {
	// A list of Authorization Scope.
	Scopes []types.Cognito_ResourceServerScope `json:"scopes,omitempty" yaml:"scopes,omitempty"`

	// User pool the client belongs to.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// An identifier for the resource server.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// A name for the resource server.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
