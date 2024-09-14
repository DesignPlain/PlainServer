package servicenetworking

type PeeredDnsDomain struct {
	// Private service connection between service and consumer network, defaults to `servicenetworking.googleapis.com`
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// The DNS domain suffix of the peered DNS domain. Make sure to suffix with a `.` (dot).
	DnsSuffix string `json:"dnsSuffix,omitempty" yaml:"dnsSuffix,omitempty"`

	// Internal name used for the peered DNS domain.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The network in the consumer project.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// The producer project number. If not provided, the provider project is used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
