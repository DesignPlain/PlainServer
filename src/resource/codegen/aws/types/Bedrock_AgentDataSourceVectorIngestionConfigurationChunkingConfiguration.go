package types

type Bedrock_AgentDataSourceVectorIngestionConfigurationChunkingConfiguration struct {
	// Option for chunking your source data, either in fixed-sized chunks or as one chunk. Valid values: `FIXED_SIZE`, `NONE`.
	ChunkingStrategy string `json:"chunkingStrategy,omitempty" yaml:"chunkingStrategy,omitempty"`

	// Configurations for when you choose fixed-size chunking. If you set the chunking_strategy as `NONE`, exclude this field. See `fixed_size_chunking_configuration` for details.
	FixedSizeChunkingConfiguration Bedrock_AgentDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration `json:"fixedSizeChunkingConfiguration,omitempty" yaml:"fixedSizeChunkingConfiguration,omitempty"`
}
