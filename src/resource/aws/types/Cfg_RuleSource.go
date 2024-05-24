package types

type Cfg_RuleSource struct {
	// Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to `CUSTOM_POLICY`. See Custom Policy Details Below.
	CustomPolicyDetails Cfg_RuleSourceCustomPolicyDetails `json:"customPolicyDetails,omitempty" yaml:"customPolicyDetails,omitempty"`

	// Indicates whether AWS or the customer owns and manages the AWS Config rule. Valid values are `AWS`, `CUSTOM_LAMBDA` or `CUSTOM_POLICY`. For more information about managed rules, see the [AWS Config Managed Rules documentation](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html). For more information about custom rules, see the [AWS Config Custom Rules documentation](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html). Custom Lambda Functions require permissions to allow the AWS Config service to invoke them, e.g., via the `aws.lambda.Permission` resource.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`

	// Provides the source and type of the event that causes AWS Config to evaluate your AWS resources. Only valid if `owner` is `CUSTOM_LAMBDA` or `CUSTOM_POLICY`. See Source Detail Below.
	SourceDetails []Cfg_RuleSourceSourceDetail `json:"sourceDetails,omitempty" yaml:"sourceDetails,omitempty"`

	// For AWS Config managed rules, a predefined identifier, e.g `IAM_PASSWORD_POLICY`. For custom Lambda rules, the identifier is the ARN of the Lambda Function, such as `arn:aws:lambda:us-east-1:123456789012:function:custom_rule_name` or the `arn` attribute of the `aws.lambda.Function` resource.
	SourceIdentifier string `json:"sourceIdentifier,omitempty" yaml:"sourceIdentifier,omitempty"`
}
