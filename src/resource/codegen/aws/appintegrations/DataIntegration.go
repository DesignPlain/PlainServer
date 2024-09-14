package appintegrations

import types "libds/aws/types"

type DataIntegration struct {
	// Specifies the name of the Data Integration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A block that defines the name of the data and how often it should be pulled from the source. The Schedule Config block is documented below.
	ScheduleConfig types.Appintegrations_DataIntegrationScheduleConfig `json:"scheduleConfig,omitempty" yaml:"scheduleConfig,omitempty"`

	// Specifies the URI of the data source. Create an AppFlow Connector Profile and reference the name of the profile in the URL. An example of this value for Salesforce is `Salesforce://AppFlow/example` where `example` is the name of the AppFlow Connector Profile.
	SourceUri string `json:"sourceUri,omitempty" yaml:"sourceUri,omitempty"`

	// Tags to apply to the Data Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the description of the Data Integration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the KMS key Amazon Resource Name (ARN) for the Data Integration.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`
}
