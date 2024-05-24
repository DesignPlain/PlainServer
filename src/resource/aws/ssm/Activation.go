package ssm

type Activation struct {
	// The description of the resource that you want to register.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// UTC timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time. This provider will only perform drift detection of its value when present in a configuration.
	ExpirationDate string `json:"expirationDate,omitempty" yaml:"expirationDate,omitempty"`

	// The IAM Role to attach to the managed instance.
	IamRole string `json:"iamRole,omitempty" yaml:"iamRole,omitempty"`

	// The default name of the registered managed instance.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The maximum number of managed instances you want to register. The default value is 1 instance.
	RegistrationLimit int `json:"registrationLimit,omitempty" yaml:"registrationLimit,omitempty"`

	// A map of tags to assign to the object. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
