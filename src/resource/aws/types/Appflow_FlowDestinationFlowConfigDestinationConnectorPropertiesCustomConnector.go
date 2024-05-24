package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector struct {
	// Custom properties that are specific to the connector when it's used as a destination in the flow. Maximum of 50 items.
	CustomProperties map[string]string `json:"customProperties,omitempty" yaml:"customProperties,omitempty"`

	// Entity specified in the custom connector as a destination in the flow.
	EntityName string `json:"entityName,omitempty" yaml:"entityName,omitempty"`

	// Settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination. See Error Handling Config for more details.
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	// Name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update, delete, or upsert.
	IdFieldNames []string `json:"idFieldNames,omitempty" yaml:"idFieldNames,omitempty"`

	// Type of write operation to be performed in the custom connector when it's used as destination. Valid values are `INSERT`, `UPSERT`, `UPDATE`, and `DELETE`.
	WriteOperationType string `json:"writeOperationType,omitempty" yaml:"writeOperationType,omitempty"`
}
