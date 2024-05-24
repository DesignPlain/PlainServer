package types

type Codedeploy_DeploymentGroupTriggerConfiguration struct {
	// The ARN of the SNS topic through which notifications are sent.
	TriggerTargetArn string `json:"triggerTargetArn,omitempty" yaml:"triggerTargetArn,omitempty"`

	// The event type or types for which notifications are triggered. Some values that are supported: `DeploymentStart`, `DeploymentSuccess`, `DeploymentFailure`, `DeploymentStop`, `DeploymentRollback`, `InstanceStart`, `InstanceSuccess`, `InstanceFailure`.  See [the CodeDeploy documentation](http://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-sns-event-notifications-create-trigger.html) for all possible values.
	TriggerEvents []string `json:"triggerEvents,omitempty" yaml:"triggerEvents,omitempty"`

	// The name of the notification trigger.
	TriggerName string `json:"triggerName,omitempty" yaml:"triggerName,omitempty"`
}
