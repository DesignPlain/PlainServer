package s3control

type AccessGrantsLocation struct {
	//
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	/*
	   The ARN of the IAM role that S3 Access Grants should use when fulfilling runtime access
	   requests to the location.
	*/
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// The default S3 URI `s3://` or the URI to a custom location, a specific bucket or prefix.
	LocationScope string `json:"locationScope,omitempty" yaml:"locationScope,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
