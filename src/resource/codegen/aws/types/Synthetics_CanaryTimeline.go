package types

type Synthetics_CanaryTimeline struct {
	// Date and time the canary was created.
	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	// Date and time the canary was most recently modified.
	LastModified string `json:"lastModified,omitempty" yaml:"lastModified,omitempty"`

	// Date and time that the canary's most recent run started.
	LastStarted string `json:"lastStarted,omitempty" yaml:"lastStarted,omitempty"`

	// Date and time that the canary's most recent run ended.
	LastStopped string `json:"lastStopped,omitempty" yaml:"lastStopped,omitempty"`
}
