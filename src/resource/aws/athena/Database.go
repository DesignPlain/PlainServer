package athena

import types "DesignSphere_Server/src/resource/aws/types"

type Database struct {
	// Key-value map of custom metadata properties for the database definition.
	Properties map[string]string `json:"properties,omitempty" yaml:"properties,omitempty"`

	// That an Amazon S3 canned ACL should be set to control ownership of stored query results. See ACL Configuration below.
	AclConfiguration types.Athena_DatabaseAclConfiguration `json:"aclConfiguration,omitempty" yaml:"aclConfiguration,omitempty"`

	// Name of S3 bucket to save the results of the query execution.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Description of the database.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. See Encryption Configuration below.
	EncryptionConfiguration types.Athena_DatabaseEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// AWS account ID that you expect to be the owner of the Amazon S3 bucket.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are -not- recoverable.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Name of the database to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
