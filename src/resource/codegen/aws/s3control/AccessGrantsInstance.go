package s3control

type AccessGrantsInstance struct {
	//
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The ARN of the AWS IAM Identity Center instance associated with the S3 Access Grants instance.
	IdentityCenterArn string `json:"identityCenterArn,omitempty" yaml:"identityCenterArn,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
