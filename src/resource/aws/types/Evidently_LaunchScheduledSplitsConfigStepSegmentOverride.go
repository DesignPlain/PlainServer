package types

type Evidently_LaunchScheduledSplitsConfigStepSegmentOverride struct {
	// Specifies a number indicating the order to use to evaluate segment overrides, if there are more than one. Segment overrides with lower numbers are evaluated first.
	EvaluationOrder int `json:"evaluationOrder,omitempty" yaml:"evaluationOrder,omitempty"`

	// The name or ARN of the segment to use.
	Segment string `json:"segment,omitempty" yaml:"segment,omitempty"`

	// The traffic allocation percentages among the feature variations to assign to this segment. This is a set of key-value pairs. The keys are variation names. The values represent the amount of traffic to allocate to that variation for this segment. This is expressed in thousandths of a percent, so a weight of 50000 represents 50%!o(MISSING)f traffic.
	Weights map[string]int `json:"weights,omitempty" yaml:"weights,omitempty"`
}
