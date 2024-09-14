package networksecurity

import types "libds/gcp/types"

type ClientTlsPolicy struct {
	/*
	   Set of label tags associated with the ClientTlsPolicy resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the client tls policy.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Name of the ClientTlsPolicy resource.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
	   Structure is documented below.
	*/
	ServerValidationCas []types.Networksecurity_ClientTlsPolicyServerValidationCa `json:"serverValidationCas,omitempty" yaml:"serverValidationCas,omitempty"`

	// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
	Sni string `json:"sni,omitempty" yaml:"sni,omitempty"`

	/*
	   Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
	   Structure is documented below.
	*/
	ClientCertificate types.Networksecurity_ClientTlsPolicyClientCertificate `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`

	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
