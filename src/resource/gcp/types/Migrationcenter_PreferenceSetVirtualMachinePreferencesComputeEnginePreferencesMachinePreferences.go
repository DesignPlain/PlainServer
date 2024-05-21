package types

type Migrationcenter_PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences struct {
	/*
	   Compute Engine machine series to consider for insights and recommendations. If empty, no restriction is applied on the machine series.
	   Structure is documented below.
	*/
	AllowedMachineSeries []Migrationcenter_PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesAllowedMachineSeries `json:"allowedMachineSeries,omitempty" yaml:"allowedMachineSeries,omitempty"`
}
