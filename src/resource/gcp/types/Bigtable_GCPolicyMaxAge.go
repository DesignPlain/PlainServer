package types

type Bigtable_GCPolicyMaxAge struct {
	// Number of days before applying GC policy.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	/*
	   Duration before applying GC policy (ex. "8h"). This is required when `days` isn't set

	   -----
	*/
	Duration string `json:"duration,omitempty" yaml:"duration,omitempty"`
}
