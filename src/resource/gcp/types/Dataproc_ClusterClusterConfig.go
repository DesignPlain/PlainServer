package types

type Dataproc_ClusterClusterConfig struct {
	// The name of the cloud storage bucket ultimately used to house the staging data for the cluster. If staging_bucket is specified, it will contain this value, otherwise it will be the auto generated name.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   The Customer managed encryption keys settings for the cluster.
	   Structure defined below.
	*/
	EncryptionConfig Dataproc_ClusterClusterConfigEncryptionConfig `json:"encryptionConfig,omitempty" yaml:"encryptionConfig,omitempty"`

	/*
	   Common config settings for resources of Google Compute Engine cluster
	   instances, applicable to all instances in the cluster. Structure defined below.
	*/
	GceClusterConfig Dataproc_ClusterClusterConfigGceClusterConfig `json:"gceClusterConfig,omitempty" yaml:"gceClusterConfig,omitempty"`

	/*
	   The Cloud Storage staging bucket used to stage files,
	   such as Hadoop jars, between client machines and the cluster.
	   Note: If you don't explicitly specify a `staging_bucket`
	   then GCP will auto create / assign one for you. However, you are not guaranteed
	   an auto generated bucket which is solely dedicated to your cluster; it may be shared
	   with other clusters in the same region/zone also choosing to use the auto generation
	   option.
	*/
	StagingBucket string `json:"stagingBucket,omitempty" yaml:"stagingBucket,omitempty"`

	/*
	   A Dataproc NodeGroup resource is a group of Dataproc cluster nodes that execute an assigned role.
	   Structure defined below.
	*/
	AuxiliaryNodeGroups []Dataproc_ClusterClusterConfigAuxiliaryNodeGroup `json:"auxiliaryNodeGroups,omitempty" yaml:"auxiliaryNodeGroups,omitempty"`

	/*
	   The config settings for software inside the cluster.
	   Structure defined below.
	*/
	SoftwareConfig Dataproc_ClusterClusterConfigSoftwareConfig `json:"softwareConfig,omitempty" yaml:"softwareConfig,omitempty"`

	/*
	   The Google Compute Engine config settings for the worker instances
	   in a cluster. Structure defined below.
	*/
	WorkerConfig Dataproc_ClusterClusterConfigWorkerConfig `json:"workerConfig,omitempty" yaml:"workerConfig,omitempty"`

	// Security related configuration. Structure defined below.
	SecurityConfig Dataproc_ClusterClusterConfigSecurityConfig `json:"securityConfig,omitempty" yaml:"securityConfig,omitempty"`

	/*
	   The config settings for port access on the cluster.
	   Structure defined below.
	*/
	EndpointConfig Dataproc_ClusterClusterConfigEndpointConfig `json:"endpointConfig,omitempty" yaml:"endpointConfig,omitempty"`

	/*
	   The config setting for metastore service with the cluster.
	   Structure defined below.
	   - - -
	*/
	MetastoreConfig Dataproc_ClusterClusterConfigMetastoreConfig `json:"metastoreConfig,omitempty" yaml:"metastoreConfig,omitempty"`

	/*
	   The autoscaling policy config associated with the cluster.
	   Note that once set, if `autoscaling_config` is the only field set in `cluster_config`, it can
	   only be removed by setting `policy_uri = ""`, rather than removing the whole block.
	   Structure defined below.
	*/
	AutoscalingConfig Dataproc_ClusterClusterConfigAutoscalingConfig `json:"autoscalingConfig,omitempty" yaml:"autoscalingConfig,omitempty"`

	/*
	   Commands to execute on each node after config is completed.
	   You can specify multiple versions of these. Structure defined below.
	*/
	InitializationActions []Dataproc_ClusterClusterConfigInitializationAction `json:"initializationActions,omitempty" yaml:"initializationActions,omitempty"`

	/*
	   The settings for auto deletion cluster schedule.
	   Structure defined below.
	*/
	LifecycleConfig Dataproc_ClusterClusterConfigLifecycleConfig `json:"lifecycleConfig,omitempty" yaml:"lifecycleConfig,omitempty"`

	/*
	   The Google Compute Engine config settings for the master instances
	   in a cluster. Structure defined below.
	*/
	MasterConfig Dataproc_ClusterClusterConfigMasterConfig `json:"masterConfig,omitempty" yaml:"masterConfig,omitempty"`

	/*
	   The Google Compute Engine config settings for the additional
	   instances in a cluster. Structure defined below.
	   - --NOTE-- : `preemptible_worker_config` is
	   an alias for the api's [secondaryWorkerConfig](https://cloud.google.com/dataproc/docs/reference/rest/v1/ClusterConfig#InstanceGroupConfig). The name doesn't necessarily mean it is preemptible and is named as
	   such for legacy/compatibility reasons.
	*/
	PreemptibleWorkerConfig Dataproc_ClusterClusterConfigPreemptibleWorkerConfig `json:"preemptibleWorkerConfig,omitempty" yaml:"preemptibleWorkerConfig,omitempty"`

	/*
	   The Cloud Storage temp bucket used to store ephemeral cluster
	   and jobs data, such as Spark and MapReduce history files.
	   Note: If you don't explicitly specify a `temp_bucket` then GCP will auto create / assign one for you.
	*/
	TempBucket string `json:"tempBucket,omitempty" yaml:"tempBucket,omitempty"`

	/*
	   The Compute Engine accelerator (GPU) configuration for these instances. Can be specified multiple times.
	   Structure defined below.
	*/
	DataprocMetricConfig Dataproc_ClusterClusterConfigDataprocMetricConfig `json:"dataprocMetricConfig,omitempty" yaml:"dataprocMetricConfig,omitempty"`
}
