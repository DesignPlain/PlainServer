package types

type Cloudrunv2_getJobTemplateTemplate struct {
	// Email address of the IAM service account associated with the Task of a Job. The service account represents the identity of the running task, and determines what permissions the task has. If not provided, the task will use the project's default service account.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout.

	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// A list of Volumes to make available to containers.
	Volumes []Cloudrunv2_getJobTemplateTemplateVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
	VpcAccesses []Cloudrunv2_getJobTemplateTemplateVpcAccess `json:"vpcAccesses,omitempty" yaml:"vpcAccesses,omitempty"`

	// Holds the single container that defines the unit of execution for this task.
	Containers []Cloudrunv2_getJobTemplateTemplateContainer `json:"containers,omitempty" yaml:"containers,omitempty"`

	// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
	EncryptionKey string `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`

	// The execution environment being used to host this Task. Possible values: ["EXECUTION_ENVIRONMENT_GEN1", "EXECUTION_ENVIRONMENT_GEN2"]
	ExecutionEnvironment string `json:"executionEnvironment,omitempty" yaml:"executionEnvironment,omitempty"`

	// Number of retries allowed per Task, before marking this Task failed.
	MaxRetries int `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
}
