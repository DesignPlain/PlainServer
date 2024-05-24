package types

type Appconfig_ExtensionActionPointAction struct {
	// Information about the action.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The action name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// An Amazon Resource Name (ARN) for an Identity and Access Management assume role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The extension URI associated to the action point in the extension definition. The URI can be an Amazon Resource Name (ARN) for one of the following: an Lambda function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
