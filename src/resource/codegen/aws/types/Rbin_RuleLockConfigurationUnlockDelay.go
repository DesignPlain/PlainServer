package types

type Rbin_RuleLockConfigurationUnlockDelay struct {
	// The unlock delay period, measured in the unit specified for UnlockDelayUnit.
	UnlockDelayValue int `json:"unlockDelayValue,omitempty" yaml:"unlockDelayValue,omitempty"`

	// The unit of time in which to measure the unlock delay. Currently, the unlock delay can be measure only in days.
	UnlockDelayUnit string `json:"unlockDelayUnit,omitempty" yaml:"unlockDelayUnit,omitempty"`
}
