package types

type Autoscalingplans_ScalingPlanApplicationSource struct {
	// ARN of a AWS CloudFormation stack.
	CloudformationStackArn string `json:"cloudformationStackArn,omitempty" yaml:"cloudformationStackArn,omitempty"`

	// Set of tags.
	TagFilters []Autoscalingplans_ScalingPlanApplicationSourceTagFilter `json:"tagFilters,omitempty" yaml:"tagFilters,omitempty"`
}
