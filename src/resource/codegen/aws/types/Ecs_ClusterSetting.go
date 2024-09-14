package types

type Ecs_ClusterSetting struct {
	// Value to assign to the setting. Valid values: `enabled`, `disabled`.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Name of the setting to manage. Valid values: `containerInsights`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
