package certificateauthority

import types "libds/gcp/types"

type CaPool struct {
	/*
	   The IssuancePolicy to control how Certificates will be issued from this CaPool.
	   Structure is documented below.
	*/
	IssuancePolicy types.Certificateauthority_CaPoolIssuancePolicy `json:"issuancePolicy,omitempty" yaml:"issuancePolicy,omitempty"`

	/*
	   Labels with user-defined metadata.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	   "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Location of the CaPool. A full list of valid locations can be found by
	   running `gcloud privateca locations list`.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name for this CaPool.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	   Structure is documented below.
	*/
	PublishingOptions types.Certificateauthority_CaPoolPublishingOptions `json:"publishingOptions,omitempty" yaml:"publishingOptions,omitempty"`

	/*
	   The Tier of this CaPool.
	   Possible values are: `ENTERPRISE`, `DEVOPS`.
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
}
