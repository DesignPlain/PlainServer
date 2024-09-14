package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoData struct {
	//
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	//
	IdFieldNames []string `json:"idFieldNames,omitempty" yaml:"idFieldNames,omitempty"`

	//
	ObjectPath string `json:"objectPath,omitempty" yaml:"objectPath,omitempty"`

	// Determines how Amazon AppFlow handles the success response that it gets from the connector after placing data. See Success Response Handling Config for more details.
	SuccessResponseHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesSapoDataSuccessResponseHandlingConfig `json:"successResponseHandlingConfig,omitempty" yaml:"successResponseHandlingConfig,omitempty"`

	//
	WriteOperationType string `json:"writeOperationType,omitempty" yaml:"writeOperationType,omitempty"`
}
