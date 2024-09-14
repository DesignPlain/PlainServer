package types

type Gkeonprem_BareMetalClusterUpgradePolicy struct {
	/*
	   Specifies which upgrade policy to use.
	   Possible values are: `SERIAL`, `CONCURRENT`.
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
