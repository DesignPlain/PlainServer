package types

type Dax_ClusterNode struct {
	//
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	//
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	//
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The port used by the configuration endpoint
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
