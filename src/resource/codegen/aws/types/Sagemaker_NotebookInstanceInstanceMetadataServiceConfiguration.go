package types

type Sagemaker_NotebookInstanceInstanceMetadataServiceConfiguration struct {
	// Indicates the minimum IMDS version that the notebook instance supports. When passed "1" is passed. This means that both IMDSv1 and IMDSv2 are supported. Valid values are `1` and `2`.
	MinimumInstanceMetadataServiceVersion string `json:"minimumInstanceMetadataServiceVersion,omitempty" yaml:"minimumInstanceMetadataServiceVersion,omitempty"`
}
