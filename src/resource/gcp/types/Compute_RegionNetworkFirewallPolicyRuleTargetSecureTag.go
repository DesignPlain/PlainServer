package types

type Compute_RegionNetworkFirewallPolicyRuleTargetSecureTag struct {
	// Name of the secure tag, created with TagManager's TagValue API. @pattern tagValues/[0-9]+
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// [Output Only] State of the secure tag, either `EFFECTIVE` or `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted or its network is deleted.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
