package sagemaker

import types "libds/aws/types"

type CodeRepository struct {
	// The name of the Code Repository (must be unique).
	CodeRepositoryName string `json:"codeRepositoryName,omitempty" yaml:"codeRepositoryName,omitempty"`

	// Specifies details about the repository. see Git Config details below.
	GitConfig types.Sagemaker_CodeRepositoryGitConfig `json:"gitConfig,omitempty" yaml:"gitConfig,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
