package types

type Ssmcontacts_PlanStageTargetChannelTargetInfo struct {
	// The number of minutes to wait before retrying to send engagement if the engagement initially failed.
	RetryIntervalInMinutes int `json:"retryIntervalInMinutes,omitempty" yaml:"retryIntervalInMinutes,omitempty"`

	// The Amazon Resource Name (ARN) of the contact channel.
	ContactChannelId string `json:"contactChannelId,omitempty" yaml:"contactChannelId,omitempty"`
}
