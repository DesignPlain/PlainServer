package chime

import types "libds/aws/types"

type SdkvoiceSipRule struct {
	// The type of trigger assigned to the SIP rule in `trigger_value`. Valid values are `RequestUriHostname` or `ToPhoneNumber`.
	TriggerType string `json:"triggerType,omitempty" yaml:"triggerType,omitempty"`

	/*
	   If `trigger_type` is `RequestUriHostname`, the value can be the outbound host name of an Amazon Chime Voice Connector. If `trigger_type` is `ToPhoneNumber`, the value can be a customer-owned phone number in the E164 format. The Sip Media Application specified in the Sip Rule is triggered if the request URI in an incoming SIP request matches the `RequestUriHostname`, or if the "To" header in the incoming SIP request matches the `ToPhoneNumber` value.

	   The following arguments are optional:
	*/
	TriggerValue string `json:"triggerValue,omitempty" yaml:"triggerValue,omitempty"`

	// Enables or disables a rule. You must disable rules before you can delete them.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// The name of the SIP rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of SIP media applications with priority and AWS Region. Only one SIP application per AWS Region can be used. See `target_applications`.
	TargetApplications []types.Chime_SdkvoiceSipRuleTargetApplication `json:"targetApplications,omitempty" yaml:"targetApplications,omitempty"`
}
