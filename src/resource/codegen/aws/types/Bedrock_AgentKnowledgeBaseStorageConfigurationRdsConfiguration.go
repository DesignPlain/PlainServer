package types

type Bedrock_AgentKnowledgeBaseStorageConfigurationRdsConfiguration struct {
	// ARN of the vector store.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Name of the table in the database.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.
	CredentialsSecretArn string `json:"credentialsSecretArn,omitempty" yaml:"credentialsSecretArn,omitempty"`

	// Name of your Amazon RDS database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Names of the fields to which to map information about the vector store. This block supports the following arguments:
	FieldMapping Bedrock_AgentKnowledgeBaseStorageConfigurationRdsConfigurationFieldMapping `json:"fieldMapping,omitempty" yaml:"fieldMapping,omitempty"`
}
