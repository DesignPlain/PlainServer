package types

type Kendra_DataSourceConfigurationS3Configuration struct {
	// A block that defines the Document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes. Each metadata file contains metadata about a single document. Detailed below.
	DocumentsMetadataConfiguration Kendra_DataSourceConfigurationS3ConfigurationDocumentsMetadataConfiguration `json:"documentsMetadataConfiguration,omitempty" yaml:"documentsMetadataConfiguration,omitempty"`

	// A list of glob patterns for documents that should not be indexed. If a document that matches an inclusion prefix or inclusion pattern also matches an exclusion pattern, the document is not indexed. Refer to [Exclusion Patterns for more examples](https://docs.aws.amazon.com/kendra/latest/dg/API_S3DataSourceConfiguration.html#Kendra-Type-S3DataSourceConfiguration-ExclusionPatterns).
	ExclusionPatterns []string `json:"exclusionPatterns,omitempty" yaml:"exclusionPatterns,omitempty"`

	// A list of glob patterns for documents that should be indexed. If a document that matches an inclusion pattern also matches an exclusion pattern, the document is not indexed. Refer to [Inclusion Patterns for more examples](https://docs.aws.amazon.com/kendra/latest/dg/API_S3DataSourceConfiguration.html#Kendra-Type-S3DataSourceConfiguration-InclusionPatterns).
	InclusionPatterns []string `json:"inclusionPatterns,omitempty" yaml:"inclusionPatterns,omitempty"`

	// A list of S3 prefixes for the documents that should be included in the index.
	InclusionPrefixes []string `json:"inclusionPrefixes,omitempty" yaml:"inclusionPrefixes,omitempty"`

	// A block that provides the path to the S3 bucket that contains the user context filtering files for the data source. For the format of the file, see [Access control for S3 data sources](https://docs.aws.amazon.com/kendra/latest/dg/s3-acl.html). Detailed below.
	AccessControlListConfiguration Kendra_DataSourceConfigurationS3ConfigurationAccessControlListConfiguration `json:"accessControlListConfiguration,omitempty" yaml:"accessControlListConfiguration,omitempty"`

	// The name of the bucket that contains the documents.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
