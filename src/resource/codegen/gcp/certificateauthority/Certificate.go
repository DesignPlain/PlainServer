package certificateauthority

import types "libds/gcp/types"

type Certificate struct {
	// The name for this Certificate.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr string `json:"pemCsr,omitempty" yaml:"pemCsr,omitempty"`

	// The name of the CaPool this Certificate belongs to.
	Pool string `json:"pool,omitempty" yaml:"pool,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Labels with user-defined metadata to apply to this resource.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	   "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	   fractional digits, terminated by 's'. Example: "3.5s".
	*/
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`

	/*
	   The config used to create a self-signed X.509 certificate or CSR.
	   Structure is documented below.
	*/
	Config types.Certificateauthority_CertificateConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   Location of the Certificate. A full list of valid locations can be found by
	   running `gcloud privateca locations list`.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The Certificate Authority ID that should issue the certificate. For example, to issue a Certificate from
	   a Certificate Authority with resource name `projects/my-project/locations/us-central1/caPools/my-pool/certificateAuthorities/my-ca`,
	   argument `pool` should be set to `projects/my-project/locations/us-central1/caPools/my-pool`, argument `certificate_authority`
	   should be set to `my-ca`.
	*/
	CertificateAuthority string `json:"certificateAuthority,omitempty" yaml:"certificateAuthority,omitempty"`

	/*
	   The resource name for a CertificateTemplate used to issue this certificate,
	   in the format `projects/-/locations/-/certificateTemplates/-`. If this is specified,
	   the caller must have the necessary permission to use this template. If this is
	   omitted, no template will be used. This template must be in the same location
	   as the Certificate.
	*/
	CertificateTemplate string `json:"certificateTemplate,omitempty" yaml:"certificateTemplate,omitempty"`
}
