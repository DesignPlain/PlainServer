package types

type Iot_ProvisioningTemplatePreProvisioningHook struct {
	// The version of the payload that was sent to the target function. The only valid (and the default) payload version is `"2020-04-01"`.
	PayloadVersion string `json:"payloadVersion,omitempty" yaml:"payloadVersion,omitempty"`

	// The ARN of the target function.
	TargetArn string `json:"targetArn,omitempty" yaml:"targetArn,omitempty"`
}
