package vmwareengine

import types "DesignSphere_Server/src/resource/gcp/types"

type PrivateCloud struct {
	// The ID of the PrivateCloud.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Network configuration in the consumer project with which the peering has to be done.
	   Structure is documented below.
	*/
	NetworkConfig types.Vmwareengine_PrivateCloudNetworkConfig `json:"networkConfig,omitempty" yaml:"networkConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Initial type of the private cloud.
	   Possible values are: `STANDARD`, `TIME_LIMITED`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// User-provided description for this private cloud.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The location where the PrivateCloud should reside.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The management cluster for this private cloud. This used for creating and managing the default cluster.
	   Structure is documented below.
	*/
	ManagementCluster types.Vmwareengine_PrivateCloudManagementCluster `json:"managementCluster,omitempty" yaml:"managementCluster,omitempty"`
}
