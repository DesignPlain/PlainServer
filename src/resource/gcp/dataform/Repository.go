package dataform

import types "DesignSphere_Server/src/resource/gcp/types"

type Repository struct {
	// A reference to the region
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Optional. The repository's user-friendly name.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Optional. Repository user labels.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format projects/-/secrets/-/versions/-. The file itself must be in a JSON format.
	NpmrcEnvironmentVariablesSecretVersion string `json:"npmrcEnvironmentVariablesSecretVersion,omitempty" yaml:"npmrcEnvironmentVariablesSecretVersion,omitempty"`

	// The service account to run workflow invocations under.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   If set, fields of workspaceCompilationOverrides override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results.
	   Structure is documented below.
	*/
	WorkspaceCompilationOverrides types.Dataform_RepositoryWorkspaceCompilationOverrides `json:"workspaceCompilationOverrides,omitempty" yaml:"workspaceCompilationOverrides,omitempty"`

	/*
	   Optional. If set, configures this repository to be linked to a Git remote.
	   Structure is documented below.
	*/
	GitRemoteSettings types.Dataform_RepositoryGitRemoteSettings `json:"gitRemoteSettings,omitempty" yaml:"gitRemoteSettings,omitempty"`

	/*
	   The repository's name.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
