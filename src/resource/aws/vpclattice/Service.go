package vpclattice

type Service struct {
	// Type of IAM policy. Either `NONE` or `AWS_IAM`.
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	// Amazon Resource Name (ARN) of the certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Custom domain name of the service.
	CustomDomainName string `json:"customDomainName,omitempty" yaml:"customDomainName,omitempty"`

	/*
	   Name of the service. The name must be unique within the account. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.Must be between 3 and 40 characters in length.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
