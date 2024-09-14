package types

type Container_getClusterMaintenancePolicyMaintenanceExclusion struct {
	//
	ExclusionName string `json:"exclusionName,omitempty" yaml:"exclusionName,omitempty"`

	// Maintenance exclusion related options.
	ExclusionOptions []Container_getClusterMaintenancePolicyMaintenanceExclusionExclusionOption `json:"exclusionOptions,omitempty" yaml:"exclusionOptions,omitempty"`

	//
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	//
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`
}
