package types

type Appflow_FlowSourceFlowConfigSourceConnectorPropertiesSalesforce struct {
	// Flag that enables dynamic fetching of new (recently added) fields in the Salesforce objects while running a flow.
	EnableDynamicFieldUpdate bool `json:"enableDynamicFieldUpdate,omitempty" yaml:"enableDynamicFieldUpdate,omitempty"`

	// Whether Amazon AppFlow includes deleted files in the flow run.
	IncludeDeletedRecords bool `json:"includeDeletedRecords,omitempty" yaml:"includeDeletedRecords,omitempty"`

	// Object specified in the flow destination.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
