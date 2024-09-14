package types

type Cloudrun_ServiceTemplateSpec struct {
	/*
	   ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
	   requests per container of the Revision. Values are:
	*/
	ContainerConcurrency int `json:"containerConcurrency,omitempty" yaml:"containerConcurrency,omitempty"`

	/*
	   Containers defines the unit of execution for this Revision.
	   Structure is documented below.
	*/
	Containers []Cloudrun_ServiceTemplateSpecContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	/*
	   Email address of the IAM service account associated with the revision of the
	   service. The service account represents the identity of the running revision,
	   and determines what permissions the revision has. If not provided, the revision
	   will use the project's default service account.
	*/
	ServiceAccountName string `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`

	/*
	   (Output, Deprecated)
	   ServingState holds a value describing the state the resources
	   are in for this Revision.
	   It is expected
	   that the system will manipulate this based on routability and load.

	   > --Warning:-- `serving_state` is deprecated and will be removed in a future major release. This field is not supported by the Cloud Run API.
	*/
	ServingState string `json:"servingState,omitempty" yaml:"servingState,omitempty"`

	// TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`

	/*
	   Volume represents a named volume in a container.
	   Structure is documented below.
	*/
	Volumes []Cloudrun_ServiceTemplateSpecVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
