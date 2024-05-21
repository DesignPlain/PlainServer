package types

type Compute_GlobalForwardingRuleServiceDirectoryRegistrations struct {
	// Service Directory namespace to register the forwarding rule under.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	/*
	   [Optional] Service Directory region to register this global forwarding rule under.
	   Default to "us-central1". Only used for PSC for Google APIs. All PSC for
	   Google APIs Forwarding Rules on the same network should use the same Service
	   Directory region.
	*/
	ServiceDirectoryRegion string `json:"serviceDirectoryRegion,omitempty" yaml:"serviceDirectoryRegion,omitempty"`
}
