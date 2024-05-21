package types

type Clouddeploy_DeliveryPipelineConditionTargetsPresentCondition struct {
	// Output only. Most recent time at which the pipeline was updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	// The list of Target names that are missing. For example, projects/{project_id}/locations/{location_name}/targets/{target_name}.
	MissingTargets []string `json:"missingTargets,omitempty" yaml:"missingTargets,omitempty"`

	// True if there aren't any missing Targets.
	Status bool `json:"status,omitempty" yaml:"status,omitempty"`
}
