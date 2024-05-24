package datasync

import types "DesignSphere_Server/src/resource/aws/types"

type LocationAzureBlob struct {
	// The SAS configuration that allows DataSync to access your Azure Blob Storage. See configuration below.
	SasConfiguration types.Datasync_LocationAzureBlobSasConfiguration `json:"sasConfiguration,omitempty" yaml:"sasConfiguration,omitempty"`

	// Path segments if you want to limit your transfer to a virtual directory in the container.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The access tier that you want your objects or files transferred into. Valid values: `HOT`, `COOL` and `ARCHIVE`. Default: `HOT`.
	AccessTier string `json:"accessTier,omitempty" yaml:"accessTier,omitempty"`

	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `json:"agentArns,omitempty" yaml:"agentArns,omitempty"`

	// The authentication method DataSync uses to access your Azure Blob Storage. Valid values: `SAS`.
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// The type of blob that you want your objects or files to be when transferring them into Azure Blob Storage. Valid values: `BLOB`. Default: `BLOB`.
	BlobType string `json:"blobType,omitempty" yaml:"blobType,omitempty"`

	// The URL of the Azure Blob Storage container involved in your transfer.
	ContainerUrl string `json:"containerUrl,omitempty" yaml:"containerUrl,omitempty"`
}
