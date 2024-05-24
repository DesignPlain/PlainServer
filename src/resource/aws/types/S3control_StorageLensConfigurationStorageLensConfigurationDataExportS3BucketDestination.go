package types

type S3control_StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestination struct {
	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Web Services organization.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// Encryption of the metrics exports in this bucket. See Encryption below for more details.
	Encryption S3control_StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryption `json:"encryption,omitempty" yaml:"encryption,omitempty"`

	// The export format. Valid values: `CSV`, `Parquet`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// The schema version of the export file. Valid values: `V_1`.
	OutputSchemaVersion string `json:"outputSchemaVersion,omitempty" yaml:"outputSchemaVersion,omitempty"`

	// The prefix of the destination bucket where the metrics export will be delivered.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
