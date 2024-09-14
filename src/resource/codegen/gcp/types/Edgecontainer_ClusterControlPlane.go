package types

type Edgecontainer_ClusterControlPlane struct {
	/*
	   Remote control plane configuration.
	   Structure is documented below.
	*/
	Remote Edgecontainer_ClusterControlPlaneRemote `json:"remote,omitempty" yaml:"remote,omitempty"`

	/*
	   Local control plane configuration.
	   Structure is documented below.
	*/
	Local Edgecontainer_ClusterControlPlaneLocal `json:"local,omitempty" yaml:"local,omitempty"`
}
