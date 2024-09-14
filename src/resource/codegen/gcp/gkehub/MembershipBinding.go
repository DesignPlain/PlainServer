package gkehub

type MembershipBinding struct {
	/*
	   Labels for this Membership binding.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Location of the membership


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The client-provided identifier of the membership binding.
	MembershipBindingId string `json:"membershipBindingId,omitempty" yaml:"membershipBindingId,omitempty"`

	// Id of the membership
	MembershipId string `json:"membershipId,omitempty" yaml:"membershipId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A Workspace resource name in the format
	   `projects/-/locations/-/scopes/-`.
	*/
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
