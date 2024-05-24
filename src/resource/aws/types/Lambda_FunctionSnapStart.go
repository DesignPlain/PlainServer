package types

type Lambda_FunctionSnapStart struct {
	// Conditions where snap start is enabled. Valid values are `PublishedVersions`.
	ApplyOn string `json:"applyOn,omitempty" yaml:"applyOn,omitempty"`

	//
	OptimizationStatus string `json:"optimizationStatus,omitempty" yaml:"optimizationStatus,omitempty"`
}
