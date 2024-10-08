package types

type Compute_BackendServiceSecuritySettings struct {
	/*
	   ClientTlsPolicy is a resource that specifies how a client should authenticate
	   connections to backends of a service. This resource itself does not affect
	   configuration unless it is attached to a backend service resource.
	*/
	ClientTlsPolicy string `json:"clientTlsPolicy,omitempty" yaml:"clientTlsPolicy,omitempty"`

	/*
	   A list of alternate names to verify the subject identity in the certificate.
	   If specified, the client will verify that the server certificate's subject
	   alt name matches one of the specified values.
	*/
	SubjectAltNames []string `json:"subjectAltNames,omitempty" yaml:"subjectAltNames,omitempty"`
}
