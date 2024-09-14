package datazone

import types "libds/aws/types"

type EnvironmentProfile struct {
	// ID of the blueprint which the environment will be created with.
	EnvironmentBlueprintIdentifier string `json:"environmentBlueprintIdentifier,omitempty" yaml:"environmentBlueprintIdentifier,omitempty"`

	// Name of the environment profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Project identifier for environment profile.

	   The following arguments are optional:
	*/
	ProjectIdentifier string `json:"projectIdentifier,omitempty" yaml:"projectIdentifier,omitempty"`

	// Array of user parameters of the environment profile with the following attributes:
	UserParameters []types.Datazone_EnvironmentProfileUserParameter `json:"userParameters,omitempty" yaml:"userParameters,omitempty"`

	// Id of the AWS account being used.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// Desired region for environment profile.
	AwsAccountRegion string `json:"awsAccountRegion,omitempty" yaml:"awsAccountRegion,omitempty"`

	// Description of environment profile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Domain Identifier for environment profile.
	DomainIdentifier string `json:"domainIdentifier,omitempty" yaml:"domainIdentifier,omitempty"`
}
