package types

type Bedrock_AgentKnowledgeBaseStorageConfigurationOpensearchServerlessConfiguration struct {
	// ARN of the OpenSearch Service vector store.
	CollectionArn string `json:"collectionArn,omitempty" yaml:"collectionArn,omitempty"`

	// The names of the fields to which to map information about the vector store. This block supports the following arguments:
	FieldMapping Bedrock_AgentKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationFieldMapping `json:"fieldMapping,omitempty" yaml:"fieldMapping,omitempty"`

	// Name of the vector store.
	VectorIndexName string `json:"vectorIndexName,omitempty" yaml:"vectorIndexName,omitempty"`
}
