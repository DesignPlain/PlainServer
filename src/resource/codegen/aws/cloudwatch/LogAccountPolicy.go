package cloudwatch

type LogAccountPolicy struct {
	// Text of the account policy. Refer to the [AWS docs](https://docs.aws.amazon.com/cli/latest/reference/logs/put-account-policy.html) for more information.
	PolicyDocument string `json:"policyDocument,omitempty" yaml:"policyDocument,omitempty"`

	// Name of the account policy.
	PolicyName string `json:"policyName,omitempty" yaml:"policyName,omitempty"`

	// Type of account policy. Either `DATA_PROTECTION_POLICY` or `SUBSCRIPTION_FILTER_POLICY`. You can have one account policy per type in an account.
	PolicyType string `json:"policyType,omitempty" yaml:"policyType,omitempty"`

	// Currently defaults to and only accepts the value: `ALL`.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// Criteria for applying a subscription filter policy to a selection of log groups. The only allowable criteria selector is `LogGroupName NOT IN []`.
	SelectionCriteria string `json:"selectionCriteria,omitempty" yaml:"selectionCriteria,omitempty"`
}
