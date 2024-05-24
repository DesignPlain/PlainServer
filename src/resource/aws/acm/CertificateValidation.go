package acm

type CertificateValidation struct {
	// ARN of the certificate that is being validated.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns []string `json:"validationRecordFqdns,omitempty" yaml:"validationRecordFqdns,omitempty"`
}
