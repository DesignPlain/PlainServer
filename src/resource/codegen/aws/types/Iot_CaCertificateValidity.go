package types

type Iot_CaCertificateValidity struct {
	// The certificate is not valid after this date.
	NotAfter string `json:"notAfter,omitempty" yaml:"notAfter,omitempty"`

	// The certificate is not valid before this date.
	NotBefore string `json:"notBefore,omitempty" yaml:"notBefore,omitempty"`
}
