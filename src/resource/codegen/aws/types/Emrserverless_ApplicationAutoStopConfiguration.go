package types

type Emrserverless_ApplicationAutoStopConfiguration struct {
	// Enables the application to automatically stop after a certain amount of time being idle. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The amount of idle time in minutes after which your application will automatically stop. Defaults to `15` minutes.
	IdleTimeoutMinutes int `json:"idleTimeoutMinutes,omitempty" yaml:"idleTimeoutMinutes,omitempty"`
}
