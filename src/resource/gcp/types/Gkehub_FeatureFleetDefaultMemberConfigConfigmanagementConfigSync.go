package types

type Gkehub_FeatureFleetDefaultMemberConfigConfigmanagementConfigSync struct {
	/*
	   Git repo configuration for the cluster
	   Structure is documented below.
	*/
	Git Gkehub_FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncGit `json:"git,omitempty" yaml:"git,omitempty"`

	/*
	   OCI repo configuration for the cluster
	   Structure is documented below.
	*/
	Oci Gkehub_FeatureFleetDefaultMemberConfigConfigmanagementConfigSyncOci `json:"oci,omitempty" yaml:"oci,omitempty"`

	// Specifies whether the Config Sync Repo is in hierarchical or unstructured mode
	SourceFormat string `json:"sourceFormat,omitempty" yaml:"sourceFormat,omitempty"`
}
