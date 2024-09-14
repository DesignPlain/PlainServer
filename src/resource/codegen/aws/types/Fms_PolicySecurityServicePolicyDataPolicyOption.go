package types

type Fms_PolicySecurityServicePolicyDataPolicyOption struct {
	// Defines the deployment model to use for the firewall policy. Documented below.
	NetworkFirewallPolicy Fms_PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy `json:"networkFirewallPolicy,omitempty" yaml:"networkFirewallPolicy,omitempty"`

	//
	ThirdPartyFirewallPolicy Fms_PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy `json:"thirdPartyFirewallPolicy,omitempty" yaml:"thirdPartyFirewallPolicy,omitempty"`
}
