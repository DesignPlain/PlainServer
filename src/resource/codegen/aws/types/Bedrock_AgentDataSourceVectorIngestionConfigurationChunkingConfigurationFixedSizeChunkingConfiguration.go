package types

type Bedrock_AgentDataSourceVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration struct {
	// Maximum number of tokens to include in a chunk.
	MaxTokens int `json:"maxTokens,omitempty" yaml:"maxTokens,omitempty"`

	// Percentage of overlap between adjacent chunks of a data source.
	OverlapPercentage int `json:"overlapPercentage,omitempty" yaml:"overlapPercentage,omitempty"`
}
