package types

type Codedeploy_DeploymentGroupAlarmConfiguration struct {
	// Indicates whether the alarm configuration is enabled. This option is useful when you want to temporarily deactivate alarm monitoring for a deployment group without having to add the same alarms again later.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from CloudWatch. The default value is `false`.
	IgnorePollAlarmFailure bool `json:"ignorePollAlarmFailure,omitempty" yaml:"ignorePollAlarmFailure,omitempty"`

	// A list of alarms configured for the deployment group. _A maximum of 10 alarms can be added to a deployment group_.
	Alarms []string `json:"alarms,omitempty" yaml:"alarms,omitempty"`
}
