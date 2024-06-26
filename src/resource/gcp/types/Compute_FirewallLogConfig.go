package types

type Compute_FirewallLogConfig struct {
	/*
	   This field denotes whether to include or exclude metadata for firewall logs.
	   Possible values are: `EXCLUDE_ALL_METADATA`, `INCLUDE_ALL_METADATA`.
	*/
	Metadata string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}
