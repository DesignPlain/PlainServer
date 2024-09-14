package appengine

import types "libds/gcp/types"

type DomainMapping struct {
	/*
	   Relative name of the domain serving the application. Example: example.com.


	   - - -
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	/*
	   Whether the domain creation should override any existing mappings for this domain.
	   By default, overrides are rejected.
	   Default value is `STRICT`.
	   Possible values are: `STRICT`, `OVERRIDE`.
	*/
	OverrideStrategy string `json:"overrideStrategy,omitempty" yaml:"overrideStrategy,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	   Structure is documented below.
	*/
	SslSettings types.Appengine_DomainMappingSslSettings `json:"sslSettings,omitempty" yaml:"sslSettings,omitempty"`
}
