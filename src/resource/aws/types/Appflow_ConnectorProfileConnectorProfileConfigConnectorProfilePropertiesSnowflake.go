package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesSnowflake struct {
	// A name for the associated Amazon S3 bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The object key for the destination bucket in which Amazon AppFlow places the files.
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// The SAPOData Private Link service name to be used for private data transfers.
	PrivateLinkServiceName string `json:"privateLinkServiceName,omitempty" yaml:"privateLinkServiceName,omitempty"`

	// AWS Region of the Snowflake account.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snowflake account. This is written in the following format: `<Database>.<Schema>.<Stage Name>`.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`

	// The name of the Snowflake warehouse.
	Warehouse string `json:"warehouse,omitempty" yaml:"warehouse,omitempty"`

	// The name of the account.
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}
