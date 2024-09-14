package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigEncryptionConfig struct {
	// The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.
	GcePdKmsKeyName string `json:"gcePdKmsKeyName,omitempty" yaml:"gcePdKmsKeyName,omitempty"`
}
