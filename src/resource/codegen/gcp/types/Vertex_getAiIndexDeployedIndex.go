package types

type Vertex_getAiIndexDeployedIndex struct {
	// The ID of the DeployedIndex in the above IndexEndpoint.
	DeployedIndexId string `json:"deployedIndexId,omitempty" yaml:"deployedIndexId,omitempty"`

	// A resource name of the IndexEndpoint.
	IndexEndpoint string `json:"indexEndpoint,omitempty" yaml:"indexEndpoint,omitempty"`
}
