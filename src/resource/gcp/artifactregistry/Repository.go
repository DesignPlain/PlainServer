package artifactregistry

import types "DesignSphere_Server/src/resource/gcp/types"

type Repository struct {
	/*
	   MavenRepositoryConfig is maven related repository details.
	   Provides additional configuration details for repositories of the maven
	   format type.
	   Structure is documented below.
	*/
	MavenConfig types.Artifactregistry_RepositoryMavenConfig `json:"mavenConfig,omitempty" yaml:"mavenConfig,omitempty"`

	/*
	   Configuration specific for a Remote Repository.
	   Structure is documented below.
	*/
	RemoteRepositoryConfig types.Artifactregistry_RepositoryRemoteRepositoryConfig `json:"remoteRepositoryConfig,omitempty" yaml:"remoteRepositoryConfig,omitempty"`

	// The name of the location this repository is located in.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The user-provided description of the repository.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The format of packages that are stored in the repository. Supported formats
	   can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	   You can only create alpha formats if you are a member of the
	   [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).


	   - - -
	*/
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	/*
	   Labels with user-defined metadata.
	   This field may contain up to 64 entries. Label keys and values may be no
	   longer than 63 characters. Label keys must begin with a lowercase letter
	   and may only contain lowercase letters, numeric characters, underscores,
	   and dashes.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The mode configures the repository to serve artifacts from different sources.
	   Default value is `STANDARD_REPOSITORY`.
	   Possible values are: `STANDARD_REPOSITORY`, `VIRTUAL_REPOSITORY`, `REMOTE_REPOSITORY`.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   The last part of the repository name, for example:
	   "repo1"
	*/
	RepositoryId string `json:"repositoryId,omitempty" yaml:"repositoryId,omitempty"`

	/*
	   If true, the cleanup pipeline is prevented from deleting versions in this
	   repository.
	*/
	CleanupPolicyDryRun bool `json:"cleanupPolicyDryRun,omitempty" yaml:"cleanupPolicyDryRun,omitempty"`

	/*
	   The Cloud KMS resource name of the customer managed encryption key that’s
	   used to encrypt the contents of the Repository. Has the form:
	   `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	   This value may not be changed after the Repository has been created.
	*/
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Docker repository config contains repository level configuration for the repositories of docker type.
	   Structure is documented below.
	*/
	DockerConfig types.Artifactregistry_RepositoryDockerConfig `json:"dockerConfig,omitempty" yaml:"dockerConfig,omitempty"`

	/*
	   Configuration specific for a Virtual Repository.
	   Structure is documented below.
	*/
	VirtualRepositoryConfig types.Artifactregistry_RepositoryVirtualRepositoryConfig `json:"virtualRepositoryConfig,omitempty" yaml:"virtualRepositoryConfig,omitempty"`

	/*
	   Cleanup policies for this repository. Cleanup policies indicate when
	   certain package versions can be automatically deleted.
	   Map keys are policy IDs supplied by users during policy creation. They must
	   unique within a repository and be under 128 characters in length.
	   Structure is documented below.
	*/
	CleanupPolicies []types.Artifactregistry_RepositoryCleanupPolicy `json:"cleanupPolicies,omitempty" yaml:"cleanupPolicies,omitempty"`
}
