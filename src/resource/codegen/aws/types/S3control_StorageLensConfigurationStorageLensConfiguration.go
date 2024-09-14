package types

type S3control_StorageLensConfigurationStorageLensConfiguration struct {
	// The account-level configurations of the S3 Storage Lens configuration. See Account Level below for more details.
	AccountLevel S3control_StorageLensConfigurationStorageLensConfigurationAccountLevel `json:"accountLevel,omitempty" yaml:"accountLevel,omitempty"`

	// The Amazon Web Services organization for the S3 Storage Lens configuration. See AWS Org below for more details.
	AwsOrg S3control_StorageLensConfigurationStorageLensConfigurationAwsOrg `json:"awsOrg,omitempty" yaml:"awsOrg,omitempty"`

	// Properties of S3 Storage Lens metrics export including the destination, schema and format. See Data Export below for more details.
	DataExport S3control_StorageLensConfigurationStorageLensConfigurationDataExport `json:"dataExport,omitempty" yaml:"dataExport,omitempty"`

	// Whether the S3 Storage Lens configuration is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// What is excluded in this configuration. Conflicts with `include`. See Exclude below for more details.
	Exclude S3control_StorageLensConfigurationStorageLensConfigurationExclude `json:"exclude,omitempty" yaml:"exclude,omitempty"`

	// What is included in this configuration. Conflicts with `exclude`. See Include below for more details.
	Include S3control_StorageLensConfigurationStorageLensConfigurationInclude `json:"include,omitempty" yaml:"include,omitempty"`
}
