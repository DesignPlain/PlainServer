package types

type Fis_ExperimentTemplateStopCondition struct {
	// ARN of the CloudWatch alarm. Required if the source is a CloudWatch alarm.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Source of the condition. One of `none`, `aws:cloudwatch:alarm`.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}
