package opensearch

import types "DesignSphere_Server/src/resource/aws/types"

type DomainSamlOptions struct {
	/*
	   Name of the domain.

	   The following arguments are optional:
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// SAML authentication options for an AWS OpenSearch Domain.
	SamlOptions types.Opensearch_DomainSamlOptionsSamlOptions `json:"samlOptions,omitempty" yaml:"samlOptions,omitempty"`
}
