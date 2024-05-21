package clouddeploy

import types "DesignSphere_Server/src/resource/gcp/types"

type Target struct {
	// Configurations for all execution that relates to this `Target`. Each `ExecutionEnvironmentUsage` value may only be used in a single configuration; using the same value multiple times is an error. When one or more configurations are specified, they must include the `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values. When no configurations are specified, execution will use the default specified in `DefaultPool`.
	ExecutionConfigs []types.Clouddeploy_TargetExecutionConfig `json:"executionConfigs,omitempty" yaml:"executionConfigs,omitempty"`

	/*
	   Name of the `Target`. Format is [a-z][a-z0-9\-]{0,62}.



	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Optional. Whether or not the `Target` requires approval.
	RequireApproval bool `json:"requireApproval,omitempty" yaml:"requireApproval,omitempty"`

	/*
	   Optional. User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// Optional. Description of the `Target`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Information specifying a GKE Cluster.
	Gke types.Clouddeploy_TargetGke `json:"gke,omitempty" yaml:"gke,omitempty"`

	/*
	   Optional. Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: - Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. - All characters must use UTF-8 encoding, and international characters are allowed. - Keys must start with a lowercase letter or international character. - Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Information specifying a multiTarget.
	MultiTarget types.Clouddeploy_TargetMultiTarget `json:"multiTarget,omitempty" yaml:"multiTarget,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Information specifying a Cloud Run deployment target.
	Run types.Clouddeploy_TargetRun `json:"run,omitempty" yaml:"run,omitempty"`

	// Information specifying an Anthos Cluster.
	AnthosCluster types.Clouddeploy_TargetAnthosCluster `json:"anthosCluster,omitempty" yaml:"anthosCluster,omitempty"`

	// Optional. The deploy parameters to use for this target.
	DeployParameters map[string]string `json:"deployParameters,omitempty" yaml:"deployParameters,omitempty"`
}
