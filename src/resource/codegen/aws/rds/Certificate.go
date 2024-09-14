package rds

type Certificate struct {
	// Certificate identifier. For example, `rds-ca-rsa4096-g1`. Refer to [AWS RDS (Relational Database) Certificate Identifier](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html#UsingWithRDS.SSL.CertificateIdentifier) for more information.
	CertificateIdentifier string `json:"certificateIdentifier,omitempty" yaml:"certificateIdentifier,omitempty"`
}
