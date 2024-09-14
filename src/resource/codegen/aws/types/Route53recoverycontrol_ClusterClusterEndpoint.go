package types

type Route53recoverycontrol_ClusterClusterEndpoint struct {
	// Cluster endpoint.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Region of the endpoint.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
