package ses

type ReceiptFilter struct {
	// The IP address or address range to filter, in CIDR notation
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	// The name of the filter
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Block or Allow
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
