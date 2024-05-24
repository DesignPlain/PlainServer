package types

type Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsInputLossSettings struct {
	// The amount of time (in milliseconds) that no input is detected. After that time, an input failover will occur.
	InputLossThresholdMsec int `json:"inputLossThresholdMsec,omitempty" yaml:"inputLossThresholdMsec,omitempty"`
}
