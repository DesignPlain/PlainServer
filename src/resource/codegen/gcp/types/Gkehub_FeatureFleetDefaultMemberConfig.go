package types

type Gkehub_FeatureFleetDefaultMemberConfig struct {
	/*
	   Policy Controller spec
	   Structure is documented below.
	*/
	Policycontroller Gkehub_FeatureFleetDefaultMemberConfigPolicycontroller `json:"policycontroller,omitempty" yaml:"policycontroller,omitempty"`

	/*
	   Config Management spec
	   Structure is documented below.
	*/
	Configmanagement Gkehub_FeatureFleetDefaultMemberConfigConfigmanagement `json:"configmanagement,omitempty" yaml:"configmanagement,omitempty"`

	/*
	   Service Mesh spec
	   Structure is documented below.
	*/
	Mesh Gkehub_FeatureFleetDefaultMemberConfigMesh `json:"mesh,omitempty" yaml:"mesh,omitempty"`
}
