package route53recoveryreadiness

type RecoveryGroup struct {
	// List of cell arns to add as nested fault domains within this recovery group
	Cells []string `json:"cells,omitempty" yaml:"cells,omitempty"`

	/*
	   A unique name describing the recovery group.

	   The following argument are optional:
	*/
	RecoveryGroupName string `json:"recoveryGroupName,omitempty" yaml:"recoveryGroupName,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
