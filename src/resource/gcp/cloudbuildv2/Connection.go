package cloudbuildv2

import types "DesignSphere_Server/src/resource/gcp/types"

type Connection struct {
	// Immutable. The resource name of the connection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Allows clients to store small amounts of arbitrary data.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   Configuration for connections to github.com.
	   Structure is documented below.
	*/
	GithubConfig types.Cloudbuildv2_ConnectionGithubConfig `json:"githubConfig,omitempty" yaml:"githubConfig,omitempty"`

	/*
	   Configuration for connections to an instance of GitHub Enterprise.
	   Structure is documented below.
	*/
	GithubEnterpriseConfig types.Cloudbuildv2_ConnectionGithubEnterpriseConfig `json:"githubEnterpriseConfig,omitempty" yaml:"githubEnterpriseConfig,omitempty"`

	/*
	   Configuration for connections to gitlab.com or an instance of GitLab Enterprise.
	   Structure is documented below.
	*/
	GitlabConfig types.Cloudbuildv2_ConnectionGitlabConfig `json:"gitlabConfig,omitempty" yaml:"gitlabConfig,omitempty"`

	/*
	   The location for the resource


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
