package types

type M2_EnvironmentHighAvailabilityConfig struct {
	// Desired number of instances for the Environment.
	DesiredCapacity int `json:"desiredCapacity,omitempty" yaml:"desiredCapacity,omitempty"`
}
