package types

type Gkehub_FeatureFleetDefaultMemberConfigMesh struct {
	/*
	   Whether to automatically manage Service Mesh
	   Possible values are: `MANAGEMENT_UNSPECIFIED`, `MANAGEMENT_AUTOMATIC`, `MANAGEMENT_MANUAL`.
	*/
	Management string `json:"management,omitempty" yaml:"management,omitempty"`
}
