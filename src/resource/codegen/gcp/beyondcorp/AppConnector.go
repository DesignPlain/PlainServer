package beyondcorp

import types "libds/gcp/types"

type AppConnector struct {
	// The region of the AppConnector.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// An arbitrary user-provided name for the AppConnector.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Resource labels to represent user provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// ID of the AppConnector.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Principal information about the Identity of the AppConnector.
	   Structure is documented below.
	*/
	PrincipalInfo types.Beyondcorp_AppConnectorPrincipalInfo `json:"principalInfo,omitempty" yaml:"principalInfo,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
