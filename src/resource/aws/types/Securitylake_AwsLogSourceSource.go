package types

type Securitylake_AwsLogSourceSource struct {
	// Specify the AWS account information where you want to enable Security Lake.
	Accounts []string `json:"accounts,omitempty" yaml:"accounts,omitempty"`

	// Specify the Regions where you want to enable Security Lake.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`

	// The name for a AWS source. This must be a Regionally unique value. Valid values: `ROUTE53`, `VPC_FLOW`, `SH_FINDINGS`, `CLOUD_TRAIL_MGMT`, `LAMBDA_EXECUTION`, `S3_DATA`.
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`

	// The version for a AWS source. This must be a Regionally unique value.
	SourceVersion string `json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`
}
