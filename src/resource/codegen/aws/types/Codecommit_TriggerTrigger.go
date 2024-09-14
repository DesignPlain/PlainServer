package types

type Codecommit_TriggerTrigger struct {
	// The branches that will be included in the trigger configuration. If no branches   are specified, the trigger will apply to all branches.
	Branches []string `json:"branches,omitempty" yaml:"branches,omitempty"`

	// Any custom data associated with the trigger that will be included in the information sent to the target of the trigger.
	CustomData string `json:"customData,omitempty" yaml:"customData,omitempty"`

	// The ARN of the resource that is the target for a trigger. For example, the ARN of a topic in Amazon Simple Notification Service (SNS).
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	// The repository events that will cause the trigger to run actions in another service, such as sending a notification through Amazon Simple Notification Service (SNS). If no events are specified, the trigger will run for all repository events. Event types include: `all`, `updateReference`, `createReference`, `deleteReference`.
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`

	// The name of the trigger.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
