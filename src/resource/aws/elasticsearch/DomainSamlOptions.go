package elasticsearch

import types "DesignSphere_Server/src/resource/aws/types"

type DomainSamlOptions struct {
	/*
	   Name of the domain.

	   The following arguments are optional:
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions types.Elasticsearch_DomainSamlOptionsSamlOptions `json:"samlOptions,omitempty" yaml:"samlOptions,omitempty"`
}
