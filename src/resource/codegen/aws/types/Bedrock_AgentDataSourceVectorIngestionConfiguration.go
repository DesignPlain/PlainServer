package types

type Bedrock_AgentDataSourceVectorIngestionConfiguration struct {
	// Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. See `chunking_configuration` block for details.
	ChunkingConfiguration Bedrock_AgentDataSourceVectorIngestionConfigurationChunkingConfiguration `json:"chunkingConfiguration,omitempty" yaml:"chunkingConfiguration,omitempty"`
}
