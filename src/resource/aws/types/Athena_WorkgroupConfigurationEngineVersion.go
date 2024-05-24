package types

type Athena_WorkgroupConfigurationEngineVersion struct {
	// The engine version on which the query runs. If `selected_engine_version` is set to `AUTO`, the effective engine version is chosen by Athena.
	EffectiveEngineVersion string `json:"effectiveEngineVersion,omitempty" yaml:"effectiveEngineVersion,omitempty"`

	// Requested engine version. Defaults to `AUTO`.
	SelectedEngineVersion string `json:"selectedEngineVersion,omitempty" yaml:"selectedEngineVersion,omitempty"`
}
