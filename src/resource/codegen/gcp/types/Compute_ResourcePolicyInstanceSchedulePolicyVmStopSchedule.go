package types

type Compute_ResourcePolicyInstanceSchedulePolicyVmStopSchedule struct {
	// Specifies the frequency for the operation, using the unix-cron format.
	Schedule string `json:"schedule,omitempty" yaml:"schedule,omitempty"`
}
