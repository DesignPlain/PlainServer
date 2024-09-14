package route53recoveryreadiness

type Cell struct {
	/*
	   Unique name describing the cell.

	   The following arguments are optional:
	*/
	CellName string `json:"cellName,omitempty" yaml:"cellName,omitempty"`

	// List of cell arns to add as nested fault domains within this cell.
	Cells []string `json:"cells,omitempty" yaml:"cells,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
