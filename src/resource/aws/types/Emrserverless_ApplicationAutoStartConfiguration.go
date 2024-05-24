package types

type Emrserverless_ApplicationAutoStartConfiguration struct {
	// Enables the application to automatically start on job submission. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
