package vmwareengine

import types "libds/gcp/types"

type Cluster struct {
	/*
	   The ID of the Cluster.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The map of cluster node types in this cluster,
	   where the key is canonical identifier of the node type (corresponds to the NodeType).
	   Structure is documented below.
	*/
	NodeTypeConfigs []types.Vmwareengine_ClusterNodeTypeConfig `json:"nodeTypeConfigs,omitempty" yaml:"nodeTypeConfigs,omitempty"`

	/*
	   The resource name of the private cloud to create a new cluster in.
	   Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	   For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
