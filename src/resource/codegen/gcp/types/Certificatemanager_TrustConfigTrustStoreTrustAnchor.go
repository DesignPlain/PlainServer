package types

type Certificatemanager_TrustConfigTrustStoreTrustAnchor struct {
	/*
	   PEM root certificate of the PKI used for validation.
	   Each certificate provided in PEM format may occupy up to 5kB.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	PemCertificate string `json:"pemCertificate,omitempty" yaml:"pemCertificate,omitempty"`
}
