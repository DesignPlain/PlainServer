package types

type Evidently_FeatureEvaluationRule struct {
	// The name for the new feature. Minimum length of `1`. Maximum length of `127`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// This value is `aws.evidently.splits` if this is an evaluation rule for a launch, and it is `aws.evidently.onlineab` if this is an evaluation rule for an experiment.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
