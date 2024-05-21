package gkehub

import types "DesignSphere_Server/src/resource/gcp/types"

type FeatureMembership struct {
	// The name of the feature
	Feature string `json:"feature,omitempty" yaml:"feature,omitempty"`

	// The location of the feature
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the membership
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`

	// The location of the membership, for example, "us-central1". Default is "global".
	MembershipLocation string `json:"membershipLocation,omitempty" yaml:"membershipLocation,omitempty"`

	// Service mesh specific spec. Structure is documented below.
	Mesh types.Gkehub_FeatureMembershipMesh `json:"mesh,omitempty" yaml:"mesh,omitempty"`

	// Policy Controller-specific spec. Structure is documented below.
	Policycontroller types.Gkehub_FeatureMembershipPolicycontroller `json:"policycontroller,omitempty" yaml:"policycontroller,omitempty"`

	// The project of the feature
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Config Management-specific spec. Structure is documented below.
	Configmanagement types.Gkehub_FeatureMembershipConfigmanagement `json:"configmanagement,omitempty" yaml:"configmanagement,omitempty"`
}
