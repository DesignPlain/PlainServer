package networksecurity

type AddressGroup struct {
	// List of items.
	Items []string `json:"items,omitempty" yaml:"items,omitempty"`

	/*
	   Set of label tags associated with the AddressGroup resource.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the gateway security policy.
	   The default value is `global`.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the AddressGroup resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the parent this address group belongs to. Format: organizations/{organization_id} or projects/{project_id}.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The type of the Address Group. Possible values are "IPV4" or "IPV6".
	   Possible values are: `IPV4`, `IPV6`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Capacity of the Address Group.
	Capacity int `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// Free-text description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
