package types

type Bedrock_CustomModelTrainingMetric struct {
	// Loss metric associated with the customization job.
	TrainingLoss float64 `json:"trainingLoss,omitempty" yaml:"trainingLoss,omitempty"`
}
