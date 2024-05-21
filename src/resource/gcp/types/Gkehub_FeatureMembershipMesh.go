package types

type Gkehub_FeatureMembershipMesh struct {
	// --DEPRECATED-- Whether to automatically manage Service Mesh control planes. Possible values: CONTROL_PLANE_MANAGEMENT_UNSPECIFIED, AUTOMATIC, MANUAL
	ControlPlane string `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`

	// Whether to automatically manage Service Mesh. Can either be `MANAGEMENT_AUTOMATIC` or `MANAGEMENT_MANUAL`.
	Management string `json:"management,omitempty" yaml:"management,omitempty"`
}
