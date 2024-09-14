package compute

import types "libds/gcp/types"

type ExternalVpnGateway struct {
	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   A list of interfaces on this external VPN gateway.
	   Structure is documented below.
	*/
	Interfaces []types.Compute_ExternalVpnGatewayInterface `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`

	/*
	   Labels for the external VPN gateway resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035.  Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z?` which means
	   the first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Indicates the redundancy type of this external VPN gateway
	   Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
	*/
	RedundancyType string `json:"redundancyType,omitempty" yaml:"redundancyType,omitempty"`
}
