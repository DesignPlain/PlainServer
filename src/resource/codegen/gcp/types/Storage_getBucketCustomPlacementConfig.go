package types

type Storage_getBucketCustomPlacementConfig struct {
	// The list of individual regions that comprise a dual-region bucket. See the docs for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.
	DataLocations []string `json:"dataLocations,omitempty" yaml:"dataLocations,omitempty"`
}
