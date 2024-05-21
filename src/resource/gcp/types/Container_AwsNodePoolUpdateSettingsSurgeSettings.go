package types

type Container_AwsNodePoolUpdateSettingsSurgeSettings struct {
	// Optional. The maximum number of nodes that can be created beyond the current size of the node pool during the update process.
	MaxSurge int `json:"maxSurge,omitempty" yaml:"maxSurge,omitempty"`

	// Optional. The maximum number of nodes that can be simultaneously unavailable during the update process. A node is considered unavailable if its status is not Ready.
	MaxUnavailable int `json:"maxUnavailable,omitempty" yaml:"maxUnavailable,omitempty"`
}
