package types

type Migrationcenter_PreferenceSetVirtualMachinePreferencesRegionPreferences struct {
	// A list of preferred regions, ordered by the most preferred region first. Set only valid Google Cloud region names. See https://cloud.google.com/compute/docs/regions-zones for available regions.
	PreferredRegions []string `json:"preferredRegions,omitempty" yaml:"preferredRegions,omitempty"`
}
