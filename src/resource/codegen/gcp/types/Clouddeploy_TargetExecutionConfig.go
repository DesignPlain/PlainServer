package types

type Clouddeploy_TargetExecutionConfig struct {
	// Optional. Cloud Storage location in which to store execution outputs. This can either be a bucket ("gs://my-bucket") or a path within a bucket ("gs://my-bucket/my-dir"). If unspecified, a default bucket located in the same region will be used.
	ArtifactStorage string `json:"artifactStorage,omitempty" yaml:"artifactStorage,omitempty"`

	// Optional. Execution timeout for a Cloud Build Execution. This must be between 10m and 24h in seconds format. If unspecified, a default timeout of 1h is used.
	ExecutionTimeout string `json:"executionTimeout,omitempty" yaml:"executionTimeout,omitempty"`

	// Optional. Google service account to use for execution. If unspecified, the project execution service account (-compute@developer.gserviceaccount.com) is used.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Required. Usages when this configuration should be applied.
	Usages []string `json:"usages,omitempty" yaml:"usages,omitempty"`

	// Optional. The resource name of the `WorkerPool`, with the format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. If this optional field is unspecified, the default Cloud Build pool will be used.
	WorkerPool string `json:"workerPool,omitempty" yaml:"workerPool,omitempty"`
}
