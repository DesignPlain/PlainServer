package types

type Gkehub_FeatureSpecClusterupgrade struct {
	// Specified if other fleet should be considered as a source of upgrades. Currently, at most one upstream fleet is allowed. The fleet name should be either fleet project number or id.
	UpstreamFleets []string `json:"upstreamFleets,omitempty" yaml:"upstreamFleets,omitempty"`

	/*
	   Configuration overrides for individual upgrades.
	   Structure is documented below.
	*/
	GkeUpgradeOverrides []Gkehub_FeatureSpecClusterupgradeGkeUpgradeOverride `json:"gkeUpgradeOverrides,omitempty" yaml:"gkeUpgradeOverrides,omitempty"`

	/*
	   Post conditions to override for the specified upgrade.
	   Structure is documented below.
	*/
	PostConditions Gkehub_FeatureSpecClusterupgradePostConditions `json:"postConditions,omitempty" yaml:"postConditions,omitempty"`
}
