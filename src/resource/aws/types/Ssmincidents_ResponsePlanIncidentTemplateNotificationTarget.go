package types

type Ssmincidents_ResponsePlanIncidentTemplateNotificationTarget struct {
	/*
	   The ARN of the Amazon SNS topic.

	   The following arguments are optional:
	*/
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`
}
