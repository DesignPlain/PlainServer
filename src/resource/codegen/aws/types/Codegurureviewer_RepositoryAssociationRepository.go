package types

type Codegurureviewer_RepositoryAssociationRepository struct {
	//
	Codecommit Codegurureviewer_RepositoryAssociationRepositoryCodecommit `json:"codecommit,omitempty" yaml:"codecommit,omitempty"`

	//
	GithubEnterpriseServer Codegurureviewer_RepositoryAssociationRepositoryGithubEnterpriseServer `json:"githubEnterpriseServer,omitempty" yaml:"githubEnterpriseServer,omitempty"`

	//
	S3Bucket Codegurureviewer_RepositoryAssociationRepositoryS3Bucket `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	//
	Bitbucket Codegurureviewer_RepositoryAssociationRepositoryBitbucket `json:"bitbucket,omitempty" yaml:"bitbucket,omitempty"`
}
