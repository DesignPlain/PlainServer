package connect

type Instance struct {
	// Specifies whether contact lens is enabled. Defaults to `true`.
	ContactLensEnabled bool `json:"contactLensEnabled,omitempty" yaml:"contactLensEnabled,omitempty"`

	// The identifier for the directory if identity_management_type is `EXISTING_DIRECTORY`.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// Specifies the identity management type attached to the instance. Allowed Values are: `SAML`, `CONNECT_MANAGED`, `EXISTING_DIRECTORY`.
	IdentityManagementType string `json:"identityManagementType,omitempty" yaml:"identityManagementType,omitempty"`

	// Specifies whether inbound calls are enabled.
	InboundCallsEnabled bool `json:"inboundCallsEnabled,omitempty" yaml:"inboundCallsEnabled,omitempty"`

	// Specifies the name of the instance. Required if `directory_id` not specified.
	InstanceAlias string `json:"instanceAlias,omitempty" yaml:"instanceAlias,omitempty"`

	// Specifies whether auto resolve best voices is enabled. Defaults to `true`.
	AutoResolveBestVoicesEnabled bool `json:"autoResolveBestVoicesEnabled,omitempty" yaml:"autoResolveBestVoicesEnabled,omitempty"`

	// Specifies whether contact flow logs are enabled. Defaults to `false`.
	ContactFlowLogsEnabled bool `json:"contactFlowLogsEnabled,omitempty" yaml:"contactFlowLogsEnabled,omitempty"`

	// Specifies whether early media for outbound calls is enabled . Defaults to `true` if outbound calls is enabled.
	EarlyMediaEnabled bool `json:"earlyMediaEnabled,omitempty" yaml:"earlyMediaEnabled,omitempty"`

	// Specifies whether multi-party calls/conference is enabled. Defaults to `false`.
	MultiPartyConferenceEnabled bool `json:"multiPartyConferenceEnabled,omitempty" yaml:"multiPartyConferenceEnabled,omitempty"`

	/*
	   Specifies whether outbound calls are enabled.
	   <!-- - `use_custom_tts_voices` - (Optional) Whether use custom tts voices is enabled. Defaults to `false` -->
	*/
	OutboundCallsEnabled bool `json:"outboundCallsEnabled,omitempty" yaml:"outboundCallsEnabled,omitempty"`
}
