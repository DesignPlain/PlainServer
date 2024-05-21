package types

type Sql_getTiersTier struct {
	// The maximum disk size of this tier in bytes.
	DiskQuota int `json:"diskQuota,omitempty" yaml:"diskQuota,omitempty"`

	// The maximum ram usage of this tier in bytes.
	Ram int `json:"ram,omitempty" yaml:"ram,omitempty"`

	// The applicable regions for this tier.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`

	// An identifier for the machine type, for example, db-custom-1-3840.
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
}
