package types

type Appflow_FlowDestinationFlowConfigDestinationConnectorPropertiesCustomerProfiles struct {
	// Unique name of the Amazon Connect Customer Profiles domain.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Object specified in the Amazon Connect Customer Profiles flow destination.
	ObjectTypeName string `json:"objectTypeName,omitempty" yaml:"objectTypeName,omitempty"`
}
