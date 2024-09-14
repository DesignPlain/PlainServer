package types

type Bedrock_AgentKnowledgeBaseStorageConfigurationRedisEnterpriseCloudConfiguration struct {
	// ARN of the secret that you created in AWS Secrets Manager that is linked to your Redis Enterprise Cloud database.
	CredentialsSecretArn string `json:"credentialsSecretArn,omitempty" yaml:"credentialsSecretArn,omitempty"`

	// Endpoint URL of the Redis Enterprise Cloud database.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// The names of the fields to which to map information about the vector store. This block supports the following arguments:
	FieldMapping Bedrock_AgentKnowledgeBaseStorageConfigurationRedisEnterpriseCloudConfigurationFieldMapping `json:"fieldMapping,omitempty" yaml:"fieldMapping,omitempty"`

	// Name of the vector index.
	VectorIndexName string `json:"vectorIndexName,omitempty" yaml:"vectorIndexName,omitempty"`
}
