package networkfirewall

import types "libds/aws/types"

type FirewallPolicy struct {
	// A friendly name of the firewall policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A friendly description of the firewall policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration types.Networkfirewall_FirewallPolicyEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy types.Networkfirewall_FirewallPolicyFirewallPolicy `json:"firewallPolicy,omitempty" yaml:"firewallPolicy,omitempty"`
}
