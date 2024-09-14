package acmpca

type Permission struct {
	// Actions that the specified AWS service principal can use. These include `IssueCertificate`, `GetCertificate`, and `ListPermissions`. Note that in order for ACM to automatically rotate certificates issued by a PCA, it must be granted permission on all 3 actions, as per the example above.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`

	// ARN of the CA that grants the permissions.
	CertificateAuthorityArn string `json:"certificateAuthorityArn,omitempty" yaml:"certificateAuthorityArn,omitempty"`

	// AWS service or identity that receives the permission. At this time, the only valid principal is `acm.amazonaws.com`.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`

	// ID of the calling account
	SourceAccount string `json:"sourceAccount,omitempty" yaml:"sourceAccount,omitempty"`
}
