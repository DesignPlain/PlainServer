package acm

import types "DesignSphere_Server/src/resource/aws/types"

type Certificate struct {
	// Certificate's PEM-formatted public key
	CertificateBody string `json:"certificateBody,omitempty" yaml:"certificateBody,omitempty"`

	/*
	   Certificate's PEM-formatted chain
	   - Creating a private CA issued certificate
	*/
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`

	// Fully qualified domain name (FQDN) in the certificate.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Configuration block used to specify information about the initial validation of each domain name. Detailed below.
	   - Importing an existing certificate
	*/
	ValidationOptions []types.Acm_CertificateValidationOption `json:"validationOptions,omitempty" yaml:"validationOptions,omitempty"`

	// ARN of an ACM PCA
	CertificateAuthorityArn string `json:"certificateAuthorityArn,omitempty" yaml:"certificateAuthorityArn,omitempty"`

	/*
	   Amount of time to start automatic renewal process before expiration.
	   Has no effect if less than 60 days.
	   Represented by either
	   a subset of [RFC 3339 duration](https://www.rfc-editor.org/rfc/rfc3339) supporting years, months, and days (e.g., `P90D`),
	   or a string such as `2160h`.
	*/
	EarlyRenewalDuration string `json:"earlyRenewalDuration,omitempty" yaml:"earlyRenewalDuration,omitempty"`

	// Specifies the algorithm of the public and private key pair that your Amazon issued certificate uses to encrypt data. See [ACM Certificate characteristics](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) for more details.
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`

	// Configuration block used to set certificate options. Detailed below.
	Options types.Acm_CertificateOptions `json:"options,omitempty" yaml:"options,omitempty"`

	// Certificate's PEM-formatted private key
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	/*
	   Set of domains that should be SANs in the issued certificate.
	   To remove all elements of a previously configured list, set this value equal to an empty list (`[]`)
	*/
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`

	// Which method to use for validation. `DNS` or `EMAIL` are valid. This parameter must not be set for certificates that were imported into ACM and then into Pulumi.
	ValidationMethod string `json:"validationMethod,omitempty" yaml:"validationMethod,omitempty"`
}
