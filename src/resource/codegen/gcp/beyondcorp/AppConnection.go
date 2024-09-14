package beyondcorp

import types "libds/gcp/types"

type AppConnection struct {
	// ID of the AppConnection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Address of the remote application endpoint for the BeyondCorp AppConnection.
	   Structure is documented below.
	*/
	ApplicationEndpoint types.Beyondcorp_AppConnectionApplicationEndpoint `json:"applicationEndpoint,omitempty" yaml:"applicationEndpoint,omitempty"`

	// An arbitrary user-provided name for the AppConnection.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Gateway used by the AppConnection.
	   Structure is documented below.
	*/
	Gateway types.Beyondcorp_AppConnectionGateway `json:"gateway,omitempty" yaml:"gateway,omitempty"`

	/*
	   Resource labels to represent user provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// List of AppConnectors that are authorised to be associated with this AppConnection
	Connectors []string `json:"connectors,omitempty" yaml:"connectors,omitempty"`

	// The region of the AppConnection.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The type of network connectivity used by the AppConnection. Refer to
	   https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
	   for a list of possible values.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
