package lakeformation

type Resource struct {
	// Role that has read/write access to the resource.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Designates an AWS Identity and Access Management (IAM) service-linked role by registering this role with the Data Catalog.
	UseServiceLinkedRole bool `json:"useServiceLinkedRole,omitempty" yaml:"useServiceLinkedRole,omitempty"`

	//
	WithFederation bool `json:"withFederation,omitempty" yaml:"withFederation,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the resource.

	   The following arguments are optional:
	*/
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	/*
	   Flag to enable AWS LakeFormation hybrid access permission mode.

	   > --NOTE:-- AWS does not support registering an S3 location with an IAM role and subsequently updating the S3 location registration to a service-linked role.
	*/
	HybridAccessEnabled bool `json:"hybridAccessEnabled,omitempty" yaml:"hybridAccessEnabled,omitempty"`
}
