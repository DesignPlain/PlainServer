package types

type Kendra_DataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration struct {
	// A prefix used to filter metadata configuration files in the AWS S3 bucket. The S3 bucket might contain multiple metadata files. Use `s3_prefix` to include only the desired metadata files.
	S3Prefix string `json:"s3Prefix,omitempty" yaml:"s3Prefix,omitempty"`
}
