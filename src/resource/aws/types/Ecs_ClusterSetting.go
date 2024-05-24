package types

type Ecs_ClusterSetting struct {
	// Name of the setting to manage. Valid values: `containerInsights`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value to assign to the setting. Valid values are `enabled` and `disabled`.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
