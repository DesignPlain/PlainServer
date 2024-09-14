package types

type Container_ClusterNodeConfigConfidentialNodes struct {
	/*
	   Enable Confidential GKE Nodes for this cluster, to
	   enforce encryption of data in-use.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
