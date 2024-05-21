package types

type Edgecontainer_ClusterControlPlaneLocal struct {
	/*
	   Only machines matching this filter will be allowed to host control
	   plane nodes. The filtering language accepts strings like "name=<name>",
	   and is documented here: [AIP-160](https://google.aip.dev/160).
	*/
	MachineFilter string `json:"machineFilter,omitempty" yaml:"machineFilter,omitempty"`

	/*
	   The number of nodes to serve as replicas of the Control Plane.
	   Only 1 and 3 are supported.
	*/
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	/*
	   Name of the Google Distributed Cloud Edge zones where this node pool
	   will be created. For example: `us-central1-edge-customer-a`.
	*/
	NodeLocation string `json:"nodeLocation,omitempty" yaml:"nodeLocation,omitempty"`

	/*
	   Policy configuration about how user applications are deployed.
	   Possible values are: `SHARED_DEPLOYMENT_POLICY_UNSPECIFIED`, `ALLOWED`, `DISALLOWED`.
	*/
	SharedDeploymentPolicy string `json:"sharedDeploymentPolicy,omitempty" yaml:"sharedDeploymentPolicy,omitempty"`
}
