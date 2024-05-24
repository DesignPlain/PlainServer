package secretsmanager

type SecretPolicy struct {
	// Valid JSON document representing a [resource policy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html). Unlike `aws.secretsmanager.Secret`, where `policy` can be set to `"{}"` to delete the policy, `"{}"` is not a valid policy since `policy` is required.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	/*
	   Secret ARN.

	   The following arguments are optional:
	*/
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" yaml:"blockPublicPolicy,omitempty"`
}
