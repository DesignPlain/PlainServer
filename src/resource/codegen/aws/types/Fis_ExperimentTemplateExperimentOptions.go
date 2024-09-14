package types

type Fis_ExperimentTemplateExperimentOptions struct {
	// Specifies the account targeting setting for experiment options. Supports `single-account` and `multi-account`.
	AccountTargeting string `json:"accountTargeting,omitempty" yaml:"accountTargeting,omitempty"`

	// Specifies the empty target resolution mode for experiment options. Supports `fail` and `skip`.
	EmptyTargetResolutionMode string `json:"emptyTargetResolutionMode,omitempty" yaml:"emptyTargetResolutionMode,omitempty"`
}
