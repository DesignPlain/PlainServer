package types

type Gkehub_FeatureMembershipConfigmanagement struct {
	// Binauthz configuration for the cluster. Structure is documented below.
	Binauthz Gkehub_FeatureMembershipConfigmanagementBinauthz `json:"binauthz,omitempty" yaml:"binauthz,omitempty"`

	// Config Sync configuration for the cluster. Structure is documented below.
	ConfigSync Gkehub_FeatureMembershipConfigmanagementConfigSync `json:"configSync,omitempty" yaml:"configSync,omitempty"`

	// Hierarchy Controller configuration for the cluster. Structure is documented below.
	HierarchyController Gkehub_FeatureMembershipConfigmanagementHierarchyController `json:"hierarchyController,omitempty" yaml:"hierarchyController,omitempty"`

	// Policy Controller configuration for the cluster. Structure is documented below.
	PolicyController Gkehub_FeatureMembershipConfigmanagementPolicyController `json:"policyController,omitempty" yaml:"policyController,omitempty"`

	// Version of ACM installed.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
