package types

type Vertex_AiIndexMetadataConfigAlgorithmConfig struct {
	/*
	   Configuration options for using the tree-AH algorithm (Shallow tree + Asymmetric Hashing).
	   Please refer to this paper for more details: https://arxiv.org/abs/1908.10396
	   Structure is documented below.
	*/
	TreeAhConfig Vertex_AiIndexMetadataConfigAlgorithmConfigTreeAhConfig `json:"treeAhConfig,omitempty" yaml:"treeAhConfig,omitempty"`

	/*
	   Configuration options for using brute force search, which simply implements the
	   standard linear search in the database for each query.
	*/
	BruteForceConfig Vertex_AiIndexMetadataConfigAlgorithmConfigBruteForceConfig `json:"bruteForceConfig,omitempty" yaml:"bruteForceConfig,omitempty"`
}
