package types

type Lambda_FunctionSnapStart struct {
	// Conditions where snap start is enabled. Valid values are `PublishedVersions`.
	ApplyOn string `json:"applyOn,omitempty" yaml:"applyOn,omitempty"`

	// Optimization status of the snap start configuration. Valid values are `On` and `Off`.
	OptimizationStatus string `json:"optimizationStatus,omitempty" yaml:"optimizationStatus,omitempty"`
}
