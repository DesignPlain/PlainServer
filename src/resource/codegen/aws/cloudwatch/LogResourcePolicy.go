package cloudwatch

type LogResourcePolicy struct {
	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`

	// Name of the resource policy.
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
}
