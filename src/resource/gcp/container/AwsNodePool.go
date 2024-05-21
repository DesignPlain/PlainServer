package container

import types "DesignSphere_Server/src/resource/gcp/types"

type AwsNodePool struct {
	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The Kubernetes version to run on this node pool (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAwsServerConfig.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// The awsCluster for the resource
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// The configuration of the node pool.
	Config types.Container_AwsNodePoolConfig `json:"config,omitempty" yaml:"config,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint types.Container_AwsNodePoolMaxPodsConstraint `json:"maxPodsConstraint,omitempty" yaml:"maxPodsConstraint,omitempty"`

	// The name of this resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Autoscaler configuration for this node pool.
	Autoscaling types.Container_AwsNodePoolAutoscaling `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`

	// The Management configuration for this node pool.
	Management types.Container_AwsNodePoolManagement `json:"management,omitempty" yaml:"management,omitempty"`

	// The subnet where the node pool node run.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// Optional. Update settings control the speed and disruption of the node pool update.
	UpdateSettings types.Container_AwsNodePoolUpdateSettings `json:"updateSettings,omitempty" yaml:"updateSettings,omitempty"`
}
