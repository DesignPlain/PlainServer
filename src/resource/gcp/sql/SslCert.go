package sql

type SslCert struct {
	/*
	   The common name to be used in the certificate to identify the
	   client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
	*/
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty"`

	/*
	   The name of the Cloud SQL instance. Changing this
	   forces a new resource to be created.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
