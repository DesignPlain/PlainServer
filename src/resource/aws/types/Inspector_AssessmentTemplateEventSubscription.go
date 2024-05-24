package types

type Inspector_AssessmentTemplateEventSubscription struct {
	// The event for which you want to receive SNS notifications. Valid values are `ASSESSMENT_RUN_STARTED`, `ASSESSMENT_RUN_COMPLETED`, `ASSESSMENT_RUN_STATE_CHANGED`, and `FINDING_REPORTED`.
	Event string `json:"event,omitempty" yaml:"event,omitempty"`

	// The ARN of the SNS topic to which notifications are sent.
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
