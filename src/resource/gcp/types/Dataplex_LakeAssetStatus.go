package types

type Dataplex_LakeAssetStatus struct {
	// Number of active assets.
	ActiveAssets int `json:"activeAssets,omitempty" yaml:"activeAssets,omitempty"`

	// Number of assets that are in process of updating the security policy on attached resources.
	SecurityPolicyApplyingAssets int `json:"securityPolicyApplyingAssets,omitempty" yaml:"securityPolicyApplyingAssets,omitempty"`

	// Output only. The time when the lake was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
