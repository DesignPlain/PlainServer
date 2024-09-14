package edgecontainer

import types "libds/gcp/types"

type NodePool struct {
	// The location of the resource.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Only machines matching this filter will be allowed to join the node pool.
	   The filtering language accepts strings like "name=<name>", and is
	   documented in more detail in [AIP-160](https://google.aip.dev/160).
	*/
	MachineFilter string `json:"machineFilter,omitempty" yaml:"machineFilter,omitempty"`

	// The resource name of the node pool.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Configuration for each node in the NodePool
	   Structure is documented below.
	*/
	NodeConfig types.Edgecontainer_NodePoolNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	// The number of nodes in the pool.
	NodeCount int `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`

	// Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: `us-central1-edge-customer-a`.
	NodeLocation string `json:"nodeLocation,omitempty" yaml:"nodeLocation,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The name of the target Distributed Cloud Edge Cluster.


	   - - -
	*/
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	/*
	   Labels associated with this resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Local disk encryption options. This field is only used when enabling CMEK support.
	   Structure is documented below.
	*/
	LocalDiskEncryption types.Edgecontainer_NodePoolLocalDiskEncryption `json:"localDiskEncryption,omitempty" yaml:"localDiskEncryption,omitempty"`
}
