package container

import types "DesignSphere_Server/src/resource/gcp/types"

type AwsCluster struct {
	// Configuration related to the cluster control plane.
	ControlPlane types.Container_AwsClusterControlPlane `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`

	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Logging configuration.
	LoggingConfig types.Container_AwsClusterLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// The number of the Fleet host project where this cluster will be registered.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// Configuration related to the cluster RBAC settings.
	Authorization types.Container_AwsClusterAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`

	// Configuration options for the Binary Authorization feature.
	BinaryAuthorization types.Container_AwsClusterBinaryAuthorization `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`

	// Fleet configuration.
	Fleet types.Container_AwsClusterFleet `json:"fleet,omitempty" yaml:"fleet,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of this resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Cluster-wide networking configuration.
	Networking types.Container_AwsClusterNetworking `json:"networking,omitempty" yaml:"networking,omitempty"`

	/*
	   Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}
