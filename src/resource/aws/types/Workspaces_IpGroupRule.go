package types

type Workspaces_IpGroupRule struct {
	// The description of the IP group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The IP address range, in CIDR notation, e.g., `10.0.0.0/16`
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
