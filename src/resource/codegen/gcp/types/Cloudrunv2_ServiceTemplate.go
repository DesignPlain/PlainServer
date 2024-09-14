package types

type Cloudrunv2_ServiceTemplate struct {
	/*
	   Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc.
	   For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	   Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected.
	   All system labels in v1 now have a corresponding field in v2 RevisionTemplate.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Sets the maximum number of requests that each serving instance can receive.
	MaxInstanceRequestConcurrency int `json:"maxInstanceRequestConcurrency,omitempty" yaml:"maxInstanceRequestConcurrency,omitempty"`

	// The unique name for the revision. If this field is omitted, it will be automatically generated based on the Service name.
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`

	/*
	   Scaling settings for this Revision.
	   Structure is documented below.
	*/
	Scaling Cloudrunv2_ServiceTemplateScaling `json:"scaling,omitempty" yaml:"scaling,omitempty"`

	// Email address of the IAM service account associated with the revision of the service. The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   Max allowed time for an instance to respond to a request.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.
	   Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected.
	   All system annotations in v1 now have a corresponding field in v2 RevisionTemplate.
	   This field follows Kubernetes annotations' namespacing, limits, and rules.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   Holds the containers that define the unit of execution for this Service.
	   Structure is documented below.
	*/
	Containers []Cloudrunv2_ServiceTemplateContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	/*
	   A list of Volumes to make available to containers.
	   Structure is documented below.
	*/
	Volumes []Cloudrunv2_ServiceTemplateVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// Enables session affinity. For more information, go to https://cloud.google.com/run/docs/configuring/session-affinity
	SessionAffinity bool `json:"sessionAffinity,omitempty" yaml:"sessionAffinity,omitempty"`

	/*
	   VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
	   Structure is documented below.
	*/
	VpcAccess Cloudrunv2_ServiceTemplateVpcAccess `json:"vpcAccess,omitempty" yaml:"vpcAccess,omitempty"`

	// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
	EncryptionKey string `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`

	/*
	   The sandbox environment to host this Revision.
	   Possible values are: `EXECUTION_ENVIRONMENT_GEN1`, `EXECUTION_ENVIRONMENT_GEN2`.
	*/
	ExecutionEnvironment string `json:"executionEnvironment,omitempty" yaml:"executionEnvironment,omitempty"`
}
