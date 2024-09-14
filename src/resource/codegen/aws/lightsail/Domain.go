package lightsail

type Domain struct {
	// The name of the Lightsail domain to manage
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
