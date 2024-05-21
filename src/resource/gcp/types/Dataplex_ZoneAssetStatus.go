package types

type Dataplex_ZoneAssetStatus struct {
	// Number of active assets.
	ActiveAssets int `json:"activeAssets,omitempty" yaml:"activeAssets,omitempty"`

	// Number of assets that are in process of updating the security policy on attached resources.
	SecurityPolicyApplyingAssets int `json:"securityPolicyApplyingAssets,omitempty" yaml:"securityPolicyApplyingAssets,omitempty"`

	// Output only. The time when the zone was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
