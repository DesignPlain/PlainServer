package ecs

type AccountSettingDefault struct {
	// Name of the account setting to set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// State of the setting.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
