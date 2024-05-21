package projects

import types "DesignSphere_Server/src/resource/gcp/types"

type ApiKey struct {
	// The resource name of the key. The name must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. In another word, the name must match the regular expression: `a-z?`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Key restrictions.
	Restrictions types.Projects_ApiKeyRestrictions `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`

	// Human-readable display name of this API key. Modifiable by user.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
