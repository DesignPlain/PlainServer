package types

type Rbin_RuleLockConfiguration struct {
	// Information about the retention rule unlock delay. See `unlock_delay` below.
	UnlockDelay Rbin_RuleLockConfigurationUnlockDelay `json:"unlockDelay,omitempty" yaml:"unlockDelay,omitempty"`
}
