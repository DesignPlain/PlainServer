package types

type Migrationcenter_PreferenceSetVirtualMachinePreferences struct {
	/*
	   Commitment plan to consider when calculating costs for virtual machine insights and recommendations. If you are unsure which value to set, a 3 year commitment plan is often a good value to start with.
	   Possible values:
	   COMMITMENT_PLAN_UNSPECIFIED
	   COMMITMENT_PLAN_NONE
	   COMMITMENT_PLAN_ONE_YEAR
	   COMMITMENT_PLAN_THREE_YEARS
	*/
	CommitmentPlan string `json:"commitmentPlan,omitempty" yaml:"commitmentPlan,omitempty"`

	/*
	   The user preferences relating to Compute Engine target platform.
	   Structure is documented below.
	*/
	ComputeEnginePreferences Migrationcenter_PreferenceSetVirtualMachinePreferencesComputeEnginePreferences `json:"computeEnginePreferences,omitempty" yaml:"computeEnginePreferences,omitempty"`

	/*
	   The user preferences relating to target regions.
	   Structure is documented below.
	*/
	RegionPreferences Migrationcenter_PreferenceSetVirtualMachinePreferencesRegionPreferences `json:"regionPreferences,omitempty" yaml:"regionPreferences,omitempty"`

	/*
	   Sizing optimization strategy specifies the preferred strategy used when extrapolating usage data to calculate insights and recommendations for a virtual machine. If you are unsure which value to set, a moderate sizing optimization strategy is often a good value to start with.
	   Possible values:
	   SIZING_OPTIMIZATION_STRATEGY_UNSPECIFIED
	   SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE
	   SIZING_OPTIMIZATION_STRATEGY_MODERATE
	   SIZING_OPTIMIZATION_STRATEGY_AGGRESSIVE
	*/
	SizingOptimizationStrategy string `json:"sizingOptimizationStrategy,omitempty" yaml:"sizingOptimizationStrategy,omitempty"`

	/*
	   Preferences concerning Sole Tenancy nodes and VMs.
	   Structure is documented below.
	*/
	SoleTenancyPreferences Migrationcenter_PreferenceSetVirtualMachinePreferencesSoleTenancyPreferences `json:"soleTenancyPreferences,omitempty" yaml:"soleTenancyPreferences,omitempty"`

	/*
	   Target product for assets using this preference set. Specify either target product or business goal, but not both.
	   Possible values:
	   COMPUTE_MIGRATION_TARGET_PRODUCT_UNSPECIFIED
	   COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE
	   COMPUTE_MIGRATION_TARGET_PRODUCT_VMWARE_ENGINE
	   COMPUTE_MIGRATION_TARGET_PRODUCT_SOLE_TENANCY
	*/
	TargetProduct string `json:"targetProduct,omitempty" yaml:"targetProduct,omitempty"`

	/*
	   The user preferences relating to Google Cloud VMware Engine target platform.
	   Structure is documented below.
	*/
	VmwareEnginePreferences Migrationcenter_PreferenceSetVirtualMachinePreferencesVmwareEnginePreferences `json:"vmwareEnginePreferences,omitempty" yaml:"vmwareEnginePreferences,omitempty"`
}
