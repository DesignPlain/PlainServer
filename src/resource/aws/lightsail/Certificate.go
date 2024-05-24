package lightsail

type Certificate struct {
	// A domain name for which the certificate should be issued.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The name of the Lightsail load balancer.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of domains that should be SANs in the issued certificate. `domain_name` attribute is automatically added as a Subject Alternative Name.
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// A map of tags to assign to the resource. To create a key-only tag, use an empty string as the value. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
