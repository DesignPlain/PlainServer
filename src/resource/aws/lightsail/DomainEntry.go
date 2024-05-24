package lightsail

type DomainEntry struct {
	// The name of the Lightsail domain in which to create the entry
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// If the entry should be an alias Defaults to `false`
	IsAlias bool `json:"isAlias,omitempty" yaml:"isAlias,omitempty"`

	// Name of the entry record
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Target of the domain entry
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// Type of record
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
