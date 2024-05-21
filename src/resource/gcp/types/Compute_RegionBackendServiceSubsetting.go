package types

type Compute_RegionBackendServiceSubsetting struct {
	/*
	   The algorithm used for subsetting.
	   Possible values are: `CONSISTENT_HASH_SUBSETTING`.
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
