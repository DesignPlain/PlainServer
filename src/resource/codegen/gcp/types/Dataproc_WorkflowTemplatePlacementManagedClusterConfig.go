package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfig struct {
	// Encryption settings for the cluster.
	EncryptionConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	// The shared Compute Engine config settings for all instances in a cluster.
	GceClusterConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfig `json:"gceClusterConfig,omitempty" yaml:"gceClusterConfig,omitempty"`

	// The Compute Engine config settings for additional worker instances in a cluster.
	MasterConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigMasterConfig `json:"masterConfig,omitempty" yaml:"masterConfig,omitempty"`

	// Security settings for the cluster.
	SecurityConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigSecurityConfig `json:"securityConfig,omitempty" yaml:"securityConfig,omitempty"`

	// Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset.
	AutoscalingConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig `json:"autoscalingConfig,omitempty" yaml:"autoscalingConfig,omitempty"`

	// Port/endpoint configuration for this cluster
	EndpointConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigEndpointConfig `json:"endpointConfig,omitempty" yaml:"endpointConfig,omitempty"`

	// The config settings for software inside the cluster.
	SoftwareConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigSoftwareConfig `json:"softwareConfig,omitempty" yaml:"softwareConfig,omitempty"`

	// A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket.
	TempBucket string `json:"tempBucket,omitempty" yaml:"tempBucket,omitempty"`

	// Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's `role` metadata to run an executable on a master or worker node, as shown below using `curl` (you can also use `wget`): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if ; then ... master specific actions ... else ... worker specific actions ... fi
	InitializationActions []Dataproc_WorkflowTemplatePlacementManagedClusterConfigInitializationAction `json:"initializationActions,omitempty" yaml:"initializationActions,omitempty"`

	// Lifecycle setting for the cluster.
	LifecycleConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigLifecycleConfig `json:"lifecycleConfig,omitempty" yaml:"lifecycleConfig,omitempty"`

	// The Compute Engine config settings for additional worker instances in a cluster.
	SecondaryWorkerConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig `json:"secondaryWorkerConfig,omitempty" yaml:"secondaryWorkerConfig,omitempty"`

	// A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see (https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)).
	StagingBucket string `json:"stagingBucket,omitempty" yaml:"stagingBucket,omitempty"`

	// The Kubernetes Engine config for Dataproc clusters deployed to Kubernetes. Setting this is considered mutually exclusive with Compute Engine-based options such as `gce_cluster_config`, `master_config`, `worker_config`, `secondary_worker_config`, and `autoscaling_config`.
	GkeClusterConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig `json:"gkeClusterConfig,omitempty" yaml:"gkeClusterConfig,omitempty"`

	// Metastore configuration.
	MetastoreConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigMetastoreConfig `json:"metastoreConfig,omitempty" yaml:"metastoreConfig,omitempty"`

	/*
	   The Compute Engine config settings for additional worker instances in a cluster.

	   - - -
	*/
	WorkerConfig Dataproc_WorkflowTemplatePlacementManagedClusterConfigWorkerConfig `json:"workerConfig,omitempty" yaml:"workerConfig,omitempty"`
}
