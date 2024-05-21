package types

type Dataproc_ClusterClusterConfigEncryptionConfig struct {
	/*
	   The Cloud KMS key name to use for PD disk encryption for
	   all instances in the cluster.

	   - - -
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
