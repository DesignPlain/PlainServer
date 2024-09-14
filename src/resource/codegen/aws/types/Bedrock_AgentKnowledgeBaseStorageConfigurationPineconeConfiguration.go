package types

type Bedrock_AgentKnowledgeBaseStorageConfigurationPineconeConfiguration struct {
	// Endpoint URL for your index management page.
	ConnectionString string `json:"connectionString,omitempty" yaml:"connectionString,omitempty"`

	// ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.
	CredentialsSecretArn string `json:"credentialsSecretArn,omitempty" yaml:"credentialsSecretArn,omitempty"`

	// The names of the fields to which to map information about the vector store. This block supports the following arguments:
	FieldMapping Bedrock_AgentKnowledgeBaseStorageConfigurationPineconeConfigurationFieldMapping `json:"fieldMapping,omitempty" yaml:"fieldMapping,omitempty"`

	// Namespace to be used to write new data to your database.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
