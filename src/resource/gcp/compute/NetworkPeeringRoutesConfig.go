package compute

type NetworkPeeringRoutesConfig struct {
	/*
	   The name of the primary network for the peering.


	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	// Name of the peering.
	Peering string `json:"peering,omitempty" yaml:"peering,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Whether to export the custom routes to the peer network.
	ExportCustomRoutes bool `json:"exportCustomRoutes,omitempty" yaml:"exportCustomRoutes,omitempty"`

	// Whether to import the custom routes to the peer network.
	ImportCustomRoutes bool `json:"importCustomRoutes,omitempty" yaml:"importCustomRoutes,omitempty"`
}
