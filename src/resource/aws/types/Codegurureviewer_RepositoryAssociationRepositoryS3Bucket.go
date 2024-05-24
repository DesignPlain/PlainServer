package types

type Codegurureviewer_RepositoryAssociationRepositoryS3Bucket struct {
	// The name of the S3 bucket used for associating a new S3 repository. Note: The name must begin with `codeguru-reviewer-`.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The name of the third party source repository.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
