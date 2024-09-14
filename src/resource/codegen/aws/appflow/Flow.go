package appflow

import types "libds/aws/types"

type Flow struct {
	// Description of the flow you want to create.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption. This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn string `json:"kmsArn,omitempty" yaml:"kmsArn,omitempty"`

	// Name of the flow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A Trigger that determine how and when the flow runs.
	TriggerConfig types.Appflow_FlowTriggerConfig `json:"triggerConfig,omitempty" yaml:"triggerConfig,omitempty"`

	// A Destination Flow Config that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigs []types.Appflow_FlowDestinationFlowConfig `json:"destinationFlowConfigs,omitempty" yaml:"destinationFlowConfigs,omitempty"`

	// A Catalog that determines the configuration that Amazon AppFlow uses when it catalogs the data that’s transferred by the associated flow. When Amazon AppFlow catalogs the data from a flow, it stores metadata in a data catalog.
	MetadataCatalogConfig types.Appflow_FlowMetadataCatalogConfig `json:"metadataCatalogConfig,omitempty" yaml:"metadataCatalogConfig,omitempty"`

	// The Source Flow Config that controls how Amazon AppFlow retrieves data from the source connector.
	SourceFlowConfig types.Appflow_FlowSourceFlowConfig `json:"sourceFlowConfig,omitempty" yaml:"sourceFlowConfig,omitempty"`

	// A Task that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks []types.Appflow_FlowTask `json:"tasks,omitempty" yaml:"tasks,omitempty"`
}
