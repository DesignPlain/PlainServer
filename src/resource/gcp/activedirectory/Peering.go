package activedirectory

type Peering struct {
	/*
	   Resource labels that can contain user-provided metadata
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// - - -
	PeeringId string `json:"peeringId,omitempty" yaml:"peeringId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The current state of this Peering.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Additional information about the current status of this peering, if available.
	StatusMessage string `json:"statusMessage,omitempty" yaml:"statusMessage,omitempty"`

	// The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" yaml:"authorizedNetwork,omitempty"`

	// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
	DomainResource string `json:"domainResource,omitempty" yaml:"domainResource,omitempty"`
}
