package sagemaker

type Image struct {
	// The description of the image.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The display name of the image. When the image is added to a domain (must be unique to the domain).
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The name of the image. Must be unique to your account.
	ImageName string `json:"imageName,omitempty" yaml:"imageName,omitempty"`

	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on your behalf.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
