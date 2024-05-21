package types

type Container_AwsClusterFleet struct {
	// The name of the managed Hub Membership resource associated to this cluster. Membership names are formatted as projects/<project-number>/locations/global/membership/<cluster-id>.
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`

	// The number of the Fleet host project where this cluster will be registered.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
