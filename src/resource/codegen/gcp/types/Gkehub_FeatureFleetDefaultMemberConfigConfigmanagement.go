package types

type Gkehub_FeatureFleetDefaultMemberConfigConfigmanagement struct {
	/*
	   ConfigSync configuration for the cluster
	   Structure is documented below.
	*/
	ConfigSync Gkehub_FeatureFleetDefaultMemberConfigConfigmanagementConfigSync `json:"configSync,omitempty" yaml:"configSync,omitempty"`

	// Version of ACM installed
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
