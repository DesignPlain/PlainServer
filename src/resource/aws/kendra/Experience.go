package kendra

import types "DesignSphere_Server/src/resource/aws/types"

type Experience struct {
	// Configuration information for your Amazon Kendra experience. The provider will only perform drift detection of its value when present in a configuration. Detailed below.
	Configuration types.Kendra_ExperienceConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// A description for your Amazon Kendra experience.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The identifier of the index for your Amazon Kendra experience.
	IndexId string `json:"indexId,omitempty" yaml:"indexId,omitempty"`

	// A name for your Amazon Kendra experience.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The Amazon Resource Name (ARN) of a role with permission to access `Query API`, `QuerySuggestions API`, `SubmitFeedback API`, and `AWS SSO` that stores your user and group information. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).

	   The following arguments are optional:
	*/
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
