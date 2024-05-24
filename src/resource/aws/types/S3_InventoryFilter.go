package types

type S3_InventoryFilter struct {
	// Prefix that an object must have to be included in the inventory results.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
