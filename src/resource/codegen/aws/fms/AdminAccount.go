package fms

type AdminAccount struct {
	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall Manager administrator account. This can be an AWS Organizations master account or a member account. Defaults to the current account. Must be configured to perform drift detection.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
