package types

type Container_AwsNodePoolUpdateSettings struct {
	// Optional. Settings for surge update.
	SurgeSettings Container_AwsNodePoolUpdateSettingsSurgeSettings `json:"surgeSettings,omitempty" yaml:"surgeSettings,omitempty"`
}
