package types

type Bigtable_GCPolicyMaxVersion struct {
	/*
	   Number of version before applying the GC policy.

	   -----
	   `gc_rules` include 2 fields:
	*/
	Number int `json:"number,omitempty" yaml:"number,omitempty"`
}
