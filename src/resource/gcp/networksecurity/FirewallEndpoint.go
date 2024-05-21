package networksecurity

type FirewallEndpoint struct {
	/*
	   A map of key/value label pairs to assign to the resource.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location (zone) of the firewall endpoint.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the firewall endpoint resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name of the parent this firewall endpoint belongs to.
	   Format: organizations/{organization_id}.


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
