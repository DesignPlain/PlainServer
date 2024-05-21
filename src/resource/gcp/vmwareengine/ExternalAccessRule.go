package vmwareengine

import types "DesignSphere_Server/src/resource/gcp/types"

type ExternalAccessRule struct {
	// A list of source ports to which the external access rule applies.
	SourcePorts []string `json:"sourcePorts,omitempty" yaml:"sourcePorts,omitempty"`

	/*
	   The action that the external access rule performs.
	   Possible values are: `ALLOW`, `DENY`.
	*/
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// User-provided description for the external access rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   If destination ranges are specified, the external access rule applies only to
	   traffic that has a destination IP address in these ranges.
	   Structure is documented below.
	*/
	DestinationIpRanges []types.Vmwareengine_ExternalAccessRuleDestinationIpRange `json:"destinationIpRanges,omitempty" yaml:"destinationIpRanges,omitempty"`

	// A list of destination ports to which the external access rule applies.
	DestinationPorts []string `json:"destinationPorts,omitempty" yaml:"destinationPorts,omitempty"`

	/*
	   If source ranges are specified, the external access rule applies only to
	   traffic that has a source IP address in these ranges.
	   Structure is documented below.
	*/
	SourceIpRanges []types.Vmwareengine_ExternalAccessRuleSourceIpRange `json:"sourceIpRanges,omitempty" yaml:"sourceIpRanges,omitempty"`

	// The IP protocol to which the external access rule applies.
	IpProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`

	// The ID of the external access rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The resource name of the network policy.
	   Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	   For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// External access rule priority, which determines the external access rule to use when multiple rules apply.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
