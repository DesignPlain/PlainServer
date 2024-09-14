package types

type Fms_PolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy struct {
	// Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
	FirewallDeploymentModel string `json:"firewallDeploymentModel,omitempty" yaml:"firewallDeploymentModel,omitempty"`
}
