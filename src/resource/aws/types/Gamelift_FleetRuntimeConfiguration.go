package types

type Gamelift_FleetRuntimeConfiguration struct {
	// Maximum amount of time (in seconds) that a game session can remain in status `ACTIVATING`.
	GameSessionActivationTimeoutSeconds int `json:"gameSessionActivationTimeoutSeconds,omitempty" yaml:"gameSessionActivationTimeoutSeconds,omitempty"`

	// Maximum number of game sessions with status `ACTIVATING` to allow on an instance simultaneously.
	MaxConcurrentGameSessionActivations int `json:"maxConcurrentGameSessionActivations,omitempty" yaml:"maxConcurrentGameSessionActivations,omitempty"`

	// Collection of server process configurations that describe which server processes to run on each instance in a fleet. See below.
	ServerProcesses []Gamelift_FleetRuntimeConfigurationServerProcess `json:"serverProcesses,omitempty" yaml:"serverProcesses,omitempty"`
}
