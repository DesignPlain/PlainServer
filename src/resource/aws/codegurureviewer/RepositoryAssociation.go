package codegurureviewer

import types "DesignSphere_Server/src/resource/aws/types"

type RepositoryAssociation struct {
	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An object describing the KMS key to asssociate. Block is documented below.
	KmsKeyDetails types.Codegurureviewer_RepositoryAssociationKmsKeyDetails `json:"kmsKeyDetails,omitempty" yaml:"kmsKeyDetails,omitempty"`

	/*
	   An object describing the repository to associate. Valid values: `bitbucket`, `codecommit`, `github_enterprise_server`, or `s3_bucket`. Block is documented below. Note: for repositories that leverage CodeStar connections (ex. `bitbucket`, `github_enterprise_server`) the connection must be in `Available` status prior to creating this resource.

	   The following arguments are optional:
	*/
	Repository types.Codegurureviewer_RepositoryAssociationRepository `json:"repository,omitempty" yaml:"repository,omitempty"`
}
