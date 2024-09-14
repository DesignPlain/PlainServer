package types

type Ec2_ManagedPrefixListEntry struct {
	// CIDR block of this entry.
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	// Description of this entry. Due to API limitations, updating only the description of an existing entry requires temporarily removing and re-adding the entry.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
