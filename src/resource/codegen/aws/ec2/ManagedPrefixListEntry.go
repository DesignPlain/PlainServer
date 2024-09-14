package ec2

type ManagedPrefixListEntry struct {
	// CIDR block of this entry.
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`

	// Description of this entry. Please note that due to API limitations, updating only the description of an entry will require recreating the entry.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the prefix list.
	PrefixListId string `json:"prefixListId,omitempty" yaml:"prefixListId,omitempty"`
}
