package elasticsearch

import types "libds/aws/types"

type DomainSamlOptions struct {
	// The SAML authentication options for an AWS Elasticsearch Domain.
	SamlOptions types.Elasticsearch_DomainSamlOptionsSamlOptions `json:"samlOptions,omitempty" yaml:"samlOptions,omitempty"`

	/*
	   Name of the domain.

	   The following arguments are optional:
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
