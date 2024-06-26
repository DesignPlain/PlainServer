package types

type Container_ClusterReleaseChannel struct {
	/*
	   The selected release channel.
	   Accepted values are:
	   - UNSPECIFIED: Not set.
	   - RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
	   - REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
	   - STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky.
	*/
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`
}
