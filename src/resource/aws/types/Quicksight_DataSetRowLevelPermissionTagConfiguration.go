package types

type Quicksight_DataSetRowLevelPermissionTagConfiguration struct {
	// The status of row-level security tags. If enabled, the status is `ENABLED`. If disabled, the status is `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// A set of rules associated with row-level security, such as the tag names and columns that they are assigned to. See tag_rules.
	TagRules []Quicksight_DataSetRowLevelPermissionTagConfigurationTagRule `json:"tagRules,omitempty" yaml:"tagRules,omitempty"`
}
