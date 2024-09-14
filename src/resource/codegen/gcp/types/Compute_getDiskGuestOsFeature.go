package types

type Compute_getDiskGuestOsFeature struct {
	/*
	   URL of the disk type resource describing which disk type to use to
	   create the disk.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
