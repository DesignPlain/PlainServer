package types

type Container_ClusterMaintenancePolicyMaintenanceExclusion struct {
	//
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`

	//
	EndTime string `json:"endTime,omitempty" yaml:"endTime,omitempty"`

	//
	ExclusionName string `json:"exclusionName,omitempty" yaml:"exclusionName,omitempty"`

	// MaintenanceExclusionOptions provides maintenance exclusion related options.
	ExclusionOptions Container_ClusterMaintenancePolicyMaintenanceExclusionExclusionOptions `json:"exclusionOptions,omitempty" yaml:"exclusionOptions,omitempty"`
}
