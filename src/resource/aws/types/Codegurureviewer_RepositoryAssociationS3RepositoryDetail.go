package types

type Codegurureviewer_RepositoryAssociationS3RepositoryDetail struct {
	//
	CodeArtifacts []Codegurureviewer_RepositoryAssociationS3RepositoryDetailCodeArtifact `json:"codeArtifacts,omitempty" yaml:"codeArtifacts,omitempty"`

	// The name of the S3 bucket used for associating a new S3 repository. Note: The name must begin with `codeguru-reviewer-`.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
