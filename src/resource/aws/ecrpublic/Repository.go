package ecrpublic

import types "DesignSphere_Server/src/resource/aws/types"

type Repository struct {
	// Catalog data configuration for the repository. See below for schema.
	CatalogData types.Ecrpublic_RepositoryCatalogData `json:"catalogData,omitempty" yaml:"catalogData,omitempty"`

	//
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// Name of the repository.
	RepositoryName string `json:"repositoryName,omitempty" yaml:"repositoryName,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
