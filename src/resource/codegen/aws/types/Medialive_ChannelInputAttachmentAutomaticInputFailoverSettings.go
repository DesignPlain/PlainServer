package types

type Medialive_ChannelInputAttachmentAutomaticInputFailoverSettings struct {
	// This clear time defines the requirement a recovered input must meet to be considered healthy. The input must have no failover conditions for this length of time. Enter a time in milliseconds. This value is particularly important if the input\_preference for the failover pair is set to PRIMARY\_INPUT\_PREFERRED, because after this time, MediaLive will switch back to the primary input.
	ErrorClearTimeMsec int `json:"errorClearTimeMsec,omitempty" yaml:"errorClearTimeMsec,omitempty"`

	// A list of failover conditions. If any of these conditions occur, MediaLive will perform a failover to the other input. See Failover Condition Block for more details.
	FailoverConditions []Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverCondition `json:"failoverConditions,omitempty" yaml:"failoverConditions,omitempty"`

	// Input preference when deciding which input to make active when a previously failed input has recovered.
	InputPreference string `json:"inputPreference,omitempty" yaml:"inputPreference,omitempty"`

	// The input ID of the secondary input in the automatic input failover pair.
	SecondaryInputId string `json:"secondaryInputId,omitempty" yaml:"secondaryInputId,omitempty"`
}
