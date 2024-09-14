package iot

type ThingPrincipalAttachment struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// The name of the thing.
	Thing string `json:"thing,omitempty" yaml:"thing,omitempty"`
}
