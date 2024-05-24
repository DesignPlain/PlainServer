package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoData struct {
	// Object path specified in the SAPOData flow destination.
	ObjectPath string `json:"objectPath,omitempty" yaml:"objectPath,omitempty"`

	// Determines how Amazon AppFlow handles the success response that it gets from the connector after placing data. See Success Response Handling Config for more details.
	SuccessResponseHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig `json:"successResponseHandlingConfig,omitempty" yaml:"successResponseHandlingConfig,omitempty"`

	// Type of write operation to be performed in the custom connector when it's used as destination. Valid values are `INSERT`, `UPSERT`, `UPDATE`, and `DELETE`.
	WriteOperationType string `json:"writeOperationType,omitempty" yaml:"writeOperationType,omitempty"`

	// Settings that determine how Amazon AppFlow handles an error when placing data in the custom connector as destination. See Error Handling Config for more details.
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	// Name of the field that Amazon AppFlow uses as an ID when performing a write operation such as update, delete, or upsert.
	IdFieldNames []string `json:"idFieldNames,omitempty" yaml:"idFieldNames,omitempty"`
}
