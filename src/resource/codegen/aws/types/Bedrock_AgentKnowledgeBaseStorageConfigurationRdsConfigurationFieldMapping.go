package types

type Bedrock_AgentKnowledgeBaseStorageConfigurationRdsConfigurationFieldMapping struct {
	// Name of the field in which Amazon Bedrock stores metadata about the vector store.
	MetadataField string `json:"metadataField,omitempty" yaml:"metadataField,omitempty"`

	// Name of the field in which Amazon Bedrock stores the ID for each entry.
	PrimaryKeyField string `json:"primaryKeyField,omitempty" yaml:"primaryKeyField,omitempty"`

	// Name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.
	TextField string `json:"textField,omitempty" yaml:"textField,omitempty"`

	// Name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
	VectorField string `json:"vectorField,omitempty" yaml:"vectorField,omitempty"`
}
