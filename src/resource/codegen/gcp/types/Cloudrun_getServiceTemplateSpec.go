package types

type Cloudrun_getServiceTemplateSpec struct {
	/*
	   Email address of the IAM service account associated with the revision of the
	   service. The service account represents the identity of the running revision,
	   and determines what permissions the revision has. If not provided, the revision
	   will use the project's default service account.
	*/
	ServiceAccountName string `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`

	/*
	   ServingState holds a value describing the state the resources
	   are in for this Revision.
	   It is expected
	   that the system will manipulate this based on routability and load.
	*/
	ServingState string `json:"servingState,omitempty" yaml:"servingState,omitempty"`

	// TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`

	// Volume represents a named volume in a container.
	Volumes []Cloudrun_getServiceTemplateSpecVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	/*
	   ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
	   requests per container of the Revision. Values are:
	   - '0' thread-safe, the system should manage the max concurrency. This is
	       the default value.
	   - '1' not-thread-safe. Single concurrency
	   - '2-N' thread-safe, max concurrency of N
	*/
	ContainerConcurrency int `json:"containerConcurrency,omitempty" yaml:"containerConcurrency,omitempty"`

	// Containers defines the unit of execution for this Revision.
	Containers []Cloudrun_getServiceTemplateSpecContainer `json:"containers,omitempty" yaml:"containers,omitempty"`
}
