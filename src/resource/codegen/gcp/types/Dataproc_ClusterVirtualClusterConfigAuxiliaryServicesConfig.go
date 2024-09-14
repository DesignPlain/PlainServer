package types

type Dataproc_ClusterVirtualClusterConfigAuxiliaryServicesConfig struct {
	// The Hive Metastore configuration for this workload.
	MetastoreConfig Dataproc_ClusterVirtualClusterConfigAuxiliaryServicesConfigMetastoreConfig `json:"metastoreConfig,omitempty" yaml:"metastoreConfig,omitempty"`

	// The Spark History Server configuration for the workload.
	SparkHistoryServerConfig Dataproc_ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty" yaml:"sparkHistoryServerConfig,omitempty"`
}
