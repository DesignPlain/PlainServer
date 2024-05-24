package types

type Sagemaker_CodeRepositoryGitConfig struct {
	// The default branch for the Git repository.
	Branch string `json:"branch,omitempty" yaml:"branch,omitempty"`

	// The URL where the Git repository is located.
	RepositoryUrl string `json:"repositoryUrl,omitempty" yaml:"repositoryUrl,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository. The secret must have a staging label of AWSCURRENT and must be in the following format: `{"username": UserName, "password": Password}`
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`
}
