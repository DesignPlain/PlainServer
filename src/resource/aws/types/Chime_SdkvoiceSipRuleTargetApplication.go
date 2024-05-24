package types

type Chime_SdkvoiceSipRuleTargetApplication struct {
	// The AWS Region of the target application.
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// Priority of the SIP media application in the target list.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The SIP media application ID.
	SipMediaApplicationId string `json:"sipMediaApplicationId,omitempty" yaml:"sipMediaApplicationId,omitempty"`
}
