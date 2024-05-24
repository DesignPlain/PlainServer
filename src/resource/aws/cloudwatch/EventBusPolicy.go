package cloudwatch

type EventBusPolicy struct {
	/*
	   The name of the event bus to set the permissions on.
	   If you omit this, the permissions are set on the `default` event bus.
	*/
	EventBusName string `json:"eventBusName,omitempty" yaml:"eventBusName,omitempty"`

	// The text of the policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
