package codeartifact

import types "libds/aws/types"

type Repository struct {
	// The account number of the AWS account that owns the domain.
	DomainOwner string `json:"domainOwner,omitempty" yaml:"domainOwner,omitempty"`

	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections types.Codeartifact_RepositoryExternalConnections `json:"externalConnections,omitempty" yaml:"externalConnections,omitempty"`

	// The name of the repository to create.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstreams []types.Codeartifact_RepositoryUpstream `json:"upstreams,omitempty" yaml:"upstreams,omitempty"`

	// The description of the repository.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The domain that contains the created repository.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
