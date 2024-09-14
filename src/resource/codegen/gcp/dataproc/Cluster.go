package dataproc

import types "libds/gcp/types"

type Cluster struct {
	/*
	   The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
	   to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The name of the cluster, unique within the project and
	   zone.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the `cluster` will exist. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The region in which the cluster and associated nodes will be created in.
	   Defaults to `global`.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Allows you to configure a virtual Dataproc on GKE cluster.
	   Structure defined below.
	*/
	VirtualClusterConfig types.Dataproc_ClusterVirtualClusterConfig `json:"virtualClusterConfig,omitempty" yaml:"virtualClusterConfig,omitempty"`

	/*
	   Allows you to configure various aspects of the cluster.
	   Structure defined below.
	*/
	ClusterConfig types.Dataproc_ClusterClusterConfig `json:"clusterConfig,omitempty" yaml:"clusterConfig,omitempty"`

	/*
	   The timeout duration which allows graceful decomissioning when you change the number of worker nodes directly through a
	   terraform apply
	*/
	GracefulDecommissionTimeout string `json:"gracefulDecommissionTimeout,omitempty" yaml:"gracefulDecommissionTimeout,omitempty"`
}
