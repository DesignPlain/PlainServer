package types

type Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverCondition struct {
	// Failover condition type-specific settings. See Failover Condition Settings for more details.
	FailoverConditionSettings Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettings `json:"failoverConditionSettings,omitempty" yaml:"failoverConditionSettings,omitempty"`
}
