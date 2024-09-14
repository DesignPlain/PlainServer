package types

type Sesv2_ConfigurationSetEventDestinationEventDestinationEventBridgeDestination struct {
	// The Amazon Resource Name (ARN) of the Amazon EventBridge bus to publish email events to. Only the default bus is supported.
	EventBusArn string `json:"eventBusArn,omitempty" yaml:"eventBusArn,omitempty"`
}
