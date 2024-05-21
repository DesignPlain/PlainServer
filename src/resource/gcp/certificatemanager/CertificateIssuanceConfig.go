package certificatemanager

import types "DesignSphere_Server/src/resource/gcp/types"

type CertificateIssuanceConfig struct {
	/*
	   It specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate.
	   Must be a number between 1-99, inclusive.
	   You must set the rotation window percentage in relation to the certificate lifetime so that certificate renewal occurs at least 7 days after
	   the certificate has been issued and at least 7 days before it expires.
	*/
	RotationWindowPercentage int `json:"rotationWindowPercentage,omitempty" yaml:"rotationWindowPercentage,omitempty"`

	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Key algorithm to use when generating the private key.
	   Possible values are: `RSA_2048`, `ECDSA_P256`.
	*/
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`

	// The Certificate Manager location. If not specified, "global" is used.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   A user-defined name of the certificate issuance config.
	   CertificateIssuanceConfig names must be unique globally.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
	   Structure is documented below.
	*/
	CertificateAuthorityConfig types.Certificatemanager_CertificateIssuanceConfigCertificateAuthorityConfig `json:"certificateAuthorityConfig,omitempty" yaml:"certificateAuthorityConfig,omitempty"`

	/*
	   'Set of label tags associated with the CertificateIssuanceConfig resource.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Lifetime of issued certificates. A duration in seconds with up to nine fractional digits, ending with 's'.
	   Example: "1814400s". Valid values are from 21 days (1814400s) to 30 days (2592000s)
	*/
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`
}
