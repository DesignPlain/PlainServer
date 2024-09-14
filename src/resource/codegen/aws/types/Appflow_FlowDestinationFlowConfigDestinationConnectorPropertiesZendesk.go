package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesZendesk struct {
	//
	ErrorHandlingConfig Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfig `json:"errorHandlingConfig,omitempty" yaml:"errorHandlingConfig,omitempty"`

	//
	IdFieldNames []string `json:"idFieldNames,omitempty" yaml:"idFieldNames,omitempty"`

	//
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	//
	WriteOperationType string `json:"writeOperationType,omitempty" yaml:"writeOperationType,omitempty"`
}
