package types

type Appflow_FlowSourceFlowConfigSourceConnectorPropertiesCustomConnector struct {
	// Entity specified in the custom connector as a destination in the flow.
	EntityName string `json:"entityName,omitempty" yaml:"entityName,omitempty"`

	// Custom properties that are specific to the connector when it's used as a destination in the flow. Maximum of 50 items.
	CustomProperties map[string]string `json:"customProperties,omitempty" yaml:"customProperties,omitempty"`
}
