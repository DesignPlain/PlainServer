package types

type Bigquery_ConnectionSparkSparkHistoryServerConfig struct {
	// Resource name of an existing Dataproc Cluster to act as a Spark History Server for the connection if the form of projects/[projectId]/regions/[region]/clusters/[cluster_name].
	DataprocCluster string `json:"dataprocCluster,omitempty" yaml:"dataprocCluster,omitempty"`
}
