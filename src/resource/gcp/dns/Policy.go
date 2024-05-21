package dns

import types "DesignSphere_Server/src/resource/gcp/types"

type Policy struct {
	/*
	   Allows networks bound to this policy to receive DNS queries sent
	   by VMs or applications over VPN connections. When enabled, a
	   virtual IP address will be allocated from each of the sub-networks
	   that are bound to this policy.
	*/
	EnableInboundForwarding bool `json:"enableInboundForwarding,omitempty" yaml:"enableInboundForwarding,omitempty"`

	/*
	   Controls whether logging is enabled for the networks bound to this policy.
	   Defaults to no logging if not set.
	*/
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	/*
	   User assigned name for this policy.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   List of network names specifying networks to which this policy is applied.
	   Structure is documented below.
	*/
	Networks []types.Dns_PolicyNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Sets an alternative name server for the associated networks.
	   When specified, all DNS queries are forwarded to a name server that you choose.
	   Names such as .internal are not available when an alternative name server is specified.
	   Structure is documented below.
	*/
	AlternativeNameServerConfig types.Dns_PolicyAlternativeNameServerConfig `json:"alternativeNameServerConfig,omitempty" yaml:"alternativeNameServerConfig,omitempty"`

	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
