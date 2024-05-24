package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSalesforce struct {
	// Settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination. See Error Handling Config for more details.
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSalesforceErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	// Name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update, delete, or upsert.
	IdFieldNames []string `json:"idFieldNames,omitempty" yaml:"idFieldNames,omitempty"`

	// Object specified in the flow destination.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	// Type of write operation to be performed in the custom connector when it's used as destination. Valid values are `INSERT`, `UPSERT`, `UPDATE`, and `DELETE`.
	WriteOperationType string `json:"writeOperationType,omitempty" yaml:"writeOperationType,omitempty"`
}
