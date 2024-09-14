package msk

type ScramSecretAssociation struct {
	// Amazon Resource Name (ARN) of the MSK cluster.
	ClusterArn string `json:"clusterArn,omitempty" yaml:"clusterArn,omitempty"`

	// List of AWS Secrets Manager secret ARNs.
	SecretArnLists []string `json:"secretArnLists,omitempty" yaml:"secretArnLists,omitempty"`
}
