package types

type Ssmcontacts_PlanStageTarget struct {
	// A configuration block for specifying information about the contact channel that Incident Manager engages. See Channel Target Info for more details.
	ChannelTargetInfo Ssmcontacts_PlanStageTargetChannelTargetInfo `json:"channelTargetInfo,omitempty" yaml:"channelTargetInfo,omitempty"`

	// A configuration block for specifying information about the contact that Incident Manager engages. See Contact Target Info for more details.
	ContactTargetInfo Ssmcontacts_PlanStageTargetContactTargetInfo `json:"contactTargetInfo,omitempty" yaml:"contactTargetInfo,omitempty"`
}
