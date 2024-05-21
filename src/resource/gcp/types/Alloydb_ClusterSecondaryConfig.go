package types

type Alloydb_ClusterSecondaryConfig struct {
	/*
	   Name of the primary cluster must be in the format
	   'projects/{project}/locations/{location}/clusters/{cluster_id}'
	*/
	PrimaryClusterName string `json:"primaryClusterName,omitempty" yaml:"primaryClusterName,omitempty"`
}
