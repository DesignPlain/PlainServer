package types

type Bedrock_AgentKnowledgeBaseKnowledgeBaseConfiguration struct {
	// Type of data that the data source is converted into for the knowledge base. Valid Values: `VECTOR`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Details about the embeddings model that'sused to convert the data source. See `vector_knowledge_base_configuration` block for details.
	VectorKnowledgeBaseConfiguration Bedrock_AgentKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfiguration `json:"vectorKnowledgeBaseConfiguration,omitempty" yaml:"vectorKnowledgeBaseConfiguration,omitempty"`
}
