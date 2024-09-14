package types

type Ecs_ServiceServiceConnectConfigurationServiceTimeout struct {
	// Amount of time in seconds for the upstream to respond with a complete response per request. A value of 0 can be set to disable perRequestTimeout. Can only be set when appProtocol isn't TCP.
	PerRequestTimeoutSeconds int `json:"perRequestTimeoutSeconds,omitempty" yaml:"perRequestTimeoutSeconds,omitempty"`

	// Amount of time in seconds a connection will stay active while idle. A value of 0 can be set to disable idleTimeout.
	IdleTimeoutSeconds int `json:"idleTimeoutSeconds,omitempty" yaml:"idleTimeoutSeconds,omitempty"`
}
