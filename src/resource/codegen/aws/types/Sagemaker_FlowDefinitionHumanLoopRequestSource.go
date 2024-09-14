package types

type Sagemaker_FlowDefinitionHumanLoopRequestSource struct {
	// Specifies whether Amazon Rekognition or Amazon Textract are used as the integration source. Valid values are: `AWS/Rekognition/DetectModerationLabels/Image/V3` and `AWS/Textract/AnalyzeDocument/Forms/V1`.
	AwsManagedHumanLoopRequestSource string `json:"awsManagedHumanLoopRequestSource,omitempty" yaml:"awsManagedHumanLoopRequestSource,omitempty"`
}
