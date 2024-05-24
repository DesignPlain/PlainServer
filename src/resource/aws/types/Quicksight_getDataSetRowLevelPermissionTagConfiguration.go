package types

type Quicksight_getDataSetRowLevelPermissionTagConfiguration struct {
	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	//
	TagRules []Quicksight_getDataSetRowLevelPermissionTagConfigurationTagRule `json:"tagRules,omitempty" yaml:"tagRules,omitempty"`
}
