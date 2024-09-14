package networkfirewall

import types "libds/aws/types"

type Firewall struct {
	// A friendly description of the firewall.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration types.Networkfirewall_FirewallEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// The Amazon Resource Name (ARN) of the VPC Firewall policy.
	FirewallPolicyArn string `json:"firewallPolicyArn,omitempty" yaml:"firewallPolicyArn,omitempty"`

	// A flag indicating whether the firewall is protected against changes to the subnet associations. Use this setting to protect against accidentally modifying the subnet associations for a firewall that is in use. Defaults to `false`.
	SubnetChangeProtection bool `json:"subnetChangeProtection,omitempty" yaml:"subnetChangeProtection,omitempty"`

	// Set of configuration blocks describing the public subnets. Each subnet must belong to a different Availability Zone in the VPC. AWS Network Firewall creates a firewall endpoint in each subnet. See Subnet Mapping below for details.
	SubnetMappings []types.Networkfirewall_FirewallSubnetMapping `json:"subnetMappings,omitempty" yaml:"subnetMappings,omitempty"`

	// A flag indicating whether the firewall is protected against deletion. Use this setting to protect against accidentally deleting a firewall that is in use. Defaults to `false`.
	DeleteProtection bool `json:"deleteProtection,omitempty" yaml:"deleteProtection,omitempty"`

	// A flag indicating whether the firewall is protected against a change to the firewall policy association. Use this setting to protect against accidentally modifying the firewall policy for a firewall that is in use. Defaults to `false`.
	FirewallPolicyChangeProtection bool `json:"firewallPolicyChangeProtection,omitempty" yaml:"firewallPolicyChangeProtection,omitempty"`

	// A friendly name of the firewall.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The unique identifier of the VPC where AWS Network Firewall should create the firewall.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
