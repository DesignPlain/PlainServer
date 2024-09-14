package types

type Appengine_FlexibleAppVersionAutomaticScalingNetworkUtilization struct {
	// Target packets sent per second.
	TargetSentPacketsPerSecond int `json:"targetSentPacketsPerSecond,omitempty" yaml:"targetSentPacketsPerSecond,omitempty"`

	// Target bytes received per second.
	TargetReceivedBytesPerSecond int `json:"targetReceivedBytesPerSecond,omitempty" yaml:"targetReceivedBytesPerSecond,omitempty"`

	// Target packets received per second.
	TargetReceivedPacketsPerSecond int `json:"targetReceivedPacketsPerSecond,omitempty" yaml:"targetReceivedPacketsPerSecond,omitempty"`

	// Target bytes sent per second.
	TargetSentBytesPerSecond int `json:"targetSentBytesPerSecond,omitempty" yaml:"targetSentBytesPerSecond,omitempty"`
}
