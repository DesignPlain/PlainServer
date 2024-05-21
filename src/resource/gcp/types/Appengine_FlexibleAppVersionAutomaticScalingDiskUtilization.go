package types

type Appengine_FlexibleAppVersionAutomaticScalingDiskUtilization struct {
	// Target bytes read per second.
	TargetReadBytesPerSecond int `json:"targetReadBytesPerSecond,omitempty" yaml:"targetReadBytesPerSecond,omitempty"`

	// Target ops read per seconds.
	TargetReadOpsPerSecond int `json:"targetReadOpsPerSecond,omitempty" yaml:"targetReadOpsPerSecond,omitempty"`

	// Target bytes written per second.
	TargetWriteBytesPerSecond int `json:"targetWriteBytesPerSecond,omitempty" yaml:"targetWriteBytesPerSecond,omitempty"`

	// Target ops written per second.
	TargetWriteOpsPerSecond int `json:"targetWriteOpsPerSecond,omitempty" yaml:"targetWriteOpsPerSecond,omitempty"`
}
