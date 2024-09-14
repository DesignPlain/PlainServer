package types

type Vertex_AiMetadataStoreState struct {
	/*
	   (Output)
	   The disk utilization of the MetadataStore in bytes.
	*/
	DiskUtilizationBytes string `json:"diskUtilizationBytes,omitempty" yaml:"diskUtilizationBytes,omitempty"`
}
