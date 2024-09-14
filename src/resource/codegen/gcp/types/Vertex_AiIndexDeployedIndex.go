package types

type Vertex_AiIndexDeployedIndex struct {
	/*
	   (Output)
	   The ID of the DeployedIndex in the above IndexEndpoint.
	*/
	DeployedIndexId string `json:"deployedIndexId,omitempty" yaml:"deployedIndexId,omitempty"`

	/*
	   (Output)
	   A resource name of the IndexEndpoint.
	*/
	IndexEndpoint string `json:"indexEndpoint,omitempty" yaml:"indexEndpoint,omitempty"`
}
