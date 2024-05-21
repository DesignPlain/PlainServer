package types

type Networkconnectivity_ServiceConnectionPolicyPscConfig struct {
	// Max number of PSC connections for this policy.
	Limit string `json:"limit,omitempty" yaml:"limit,omitempty"`

	// IDs of the subnetworks or fully qualified identifiers for the subnetworks
	Subnetworks []string `json:"subnetworks,omitempty" yaml:"subnetworks,omitempty"`
}
