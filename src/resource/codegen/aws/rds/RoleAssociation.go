package rds

type RoleAssociation struct {
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// DB Instance Identifier to associate with the IAM Role.
	DbInstanceIdentifier string `json:"dbInstanceIdentifier,omitempty" yaml:"dbInstanceIdentifier,omitempty"`

	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName string `json:"featureName,omitempty" yaml:"featureName,omitempty"`
}
