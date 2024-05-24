package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesEventBridge struct {
	// Settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination. See Error Handling Config for more details.
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesEventBridgeErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	// Object specified in the flow destination.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
