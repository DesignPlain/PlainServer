package types

type Gkehub_FeatureSpecClusterupgradeGkeUpgradeOverridePostConditions struct {
	// Amount of time to "soak" after a rollout has been finished before marking it COMPLETE. Cannot exceed 30 days.
	Soaking string `json:"soaking,omitempty" yaml:"soaking,omitempty"`
}
