package types

type Ssmcontacts_getPlanStageTarget struct {
	//
	ChannelTargetInfos []Ssmcontacts_getPlanStageTargetChannelTargetInfo `json:"channelTargetInfos,omitempty" yaml:"channelTargetInfos,omitempty"`

	//
	ContactTargetInfos []Ssmcontacts_getPlanStageTargetContactTargetInfo `json:"contactTargetInfos,omitempty" yaml:"contactTargetInfos,omitempty"`
}
