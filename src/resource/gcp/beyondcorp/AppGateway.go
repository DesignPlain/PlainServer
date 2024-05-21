package beyondcorp

type AppGateway struct {
	// The region of the AppGateway.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The type of network connectivity used by the AppGateway.
	   Default value is `TYPE_UNSPECIFIED`.
	   Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// An arbitrary user-provided name for the AppGateway.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The type of hosting used by the AppGateway.
	   Default value is `HOST_TYPE_UNSPECIFIED`.
	   Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
	*/
	HostType string `json:"hostType,omitempty" yaml:"hostType,omitempty"`

	/*
	   Resource labels to represent user provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   ID of the AppGateway.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
