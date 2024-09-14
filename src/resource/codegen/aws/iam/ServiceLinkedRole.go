package iam

type ServiceLinkedRole struct {
	// The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
	AwsServiceName string `json:"awsServiceName,omitempty" yaml:"awsServiceName,omitempty"`

	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	CustomSuffix string `json:"customSuffix,omitempty" yaml:"customSuffix,omitempty"`

	// The description of the role.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Key-value mapping of tags for the IAM role. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
