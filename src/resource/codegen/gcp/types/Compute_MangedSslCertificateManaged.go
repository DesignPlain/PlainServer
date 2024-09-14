package types

type Compute_MangedSslCertificateManaged struct {
	/*
	   Domains for which a managed SSL certificate will be valid.  Currently,
	   there can be up to 100 domains in this list.
	*/
	Domains []string `json:"domains,omitempty" yaml:"domains,omitempty"`
}
