package types

type Gkehub_FeatureSpecClusterupgradeGkeUpgradeOverride struct {
	/*
	   Post conditions to override for the specified upgrade.
	   Structure is documented below.
	*/
	PostConditions Gkehub_FeatureSpecClusterupgradeGkeUpgradeOverridePostConditions `json:"postConditions,omitempty" yaml:"postConditions,omitempty"`

	/*
	   Which upgrade to override.
	   Structure is documented below.
	*/
	Upgrade Gkehub_FeatureSpecClusterupgradeGkeUpgradeOverrideUpgrade `json:"upgrade,omitempty" yaml:"upgrade,omitempty"`
}
