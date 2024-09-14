package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnector struct {
	//
	CustomProperties map[string]string `json:"customProperties,omitempty" yaml:"customProperties,omitempty"`

	//
	EntityName string `json:"entityName,omitempty" yaml:"entityName,omitempty"`

	//
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	//
	IdFieldNames []string `json:"idFieldNames,omitempty" yaml:"idFieldNames,omitempty"`

	//
	WriteOperationType string `json:"writeOperationType,omitempty" yaml:"writeOperationType,omitempty"`
}
