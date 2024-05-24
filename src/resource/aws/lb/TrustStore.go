package lb

type TrustStore struct {
	// S3 object key holding the client certificate CA bundle.
	CaCertificatesBundleS3Key string `json:"caCertificatesBundleS3Key,omitempty" yaml:"caCertificatesBundleS3Key,omitempty"`

	// Version Id of CA bundle S3 bucket object, if versioned, defaults to latest if omitted.
	CaCertificatesBundleS3ObjectVersion string `json:"caCertificatesBundleS3ObjectVersion,omitempty" yaml:"caCertificatesBundleS3ObjectVersion,omitempty"`

	// Name of the Trust Store. If omitted, the provider will assign a random, unique name. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// S3 Bucket name holding the client certificate CA bundle.
	CaCertificatesBundleS3Bucket string `json:"caCertificatesBundleS3Bucket,omitempty" yaml:"caCertificatesBundleS3Bucket,omitempty"`
}
