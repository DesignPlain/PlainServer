package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleCreateRule struct {
	//
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	//
	IntervalUnit string `json:"intervalUnit,omitempty" yaml:"intervalUnit,omitempty"`

	// Specifies the destination for snapshots created by the policy. To create snapshots in the same Region as the source resource, specify `CLOUD`. To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL`. If you omit this parameter, `CLOUD` is used by default. If the policy targets resources in an AWS Region, then you must create snapshots in the same Region as the source resource. If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost as the source resource, or in the Region of that Outpost. Valid values are `CLOUD` and `OUTPOST_LOCAL`.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// A list of times in 24 hour clock format that sets when the lifecycle policy should be evaluated. Max of 1. Conflicts with `cron_expression`. Must be set if `interval` is set.
	Times string `json:"times,omitempty" yaml:"times,omitempty"`

	// The schedule, as a Cron expression. The schedule interval must be between 1 hour and 1 year. Conflicts with `interval`, `interval_unit`, and `times`.
	CronExpression string `json:"cronExpression,omitempty" yaml:"cronExpression,omitempty"`
}
