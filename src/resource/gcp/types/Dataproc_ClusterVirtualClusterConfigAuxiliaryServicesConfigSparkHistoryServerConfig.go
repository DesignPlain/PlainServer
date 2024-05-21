package types

type Dataproc_ClusterVirtualClusterConfigAuxiliaryServicesConfigSparkHistoryServerConfig struct {
	/*
	   Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload.
	   - - -
	*/
	DataprocCluster string `json:"dataprocCluster,omitempty" yaml:"dataprocCluster,omitempty"`
}
