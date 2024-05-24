package costexplorer

type CostAllocationTag struct {
	// The status of a cost allocation tag. Valid values are `Active` and `Inactive`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The key for the cost allocation tag.
	TagKey string `json:"tagKey,omitempty" yaml:"tagKey,omitempty"`
}
