package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationRunConfigurationFlinkRunConfiguration struct {
	// When restoring from a snapshot, specifies whether the runtime is allowed to skip a state that cannot be mapped to the new program. Default is `false`.
	AllowNonRestoredState bool `json:"allowNonRestoredState,omitempty" yaml:"allowNonRestoredState,omitempty"`
}
