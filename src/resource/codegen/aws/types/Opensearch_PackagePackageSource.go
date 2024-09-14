package types

type Opensearch_PackagePackageSource struct {
	// The name of the Amazon S3 bucket containing the package.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Key (file name) of the package.
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`
}
