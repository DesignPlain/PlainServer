package apigateway

import types "DesignSphere_Server/src/resource/aws/types"

type DomainName struct {
	// ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
	OwnershipVerificationCertificateArn string `json:"ownershipVerificationCertificateArn,omitempty" yaml:"ownershipVerificationCertificateArn,omitempty"`

	// User-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
	RegionalCertificateName string `json:"regionalCertificateName,omitempty" yaml:"regionalCertificateName,omitempty"`

	// ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Certificate for the CA that issued the certificate, along with any intermediate CA certificates required to create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`

	// Private key associated with the domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
	CertificatePrivateKey string `json:"certificatePrivateKey,omitempty" yaml:"certificatePrivateKey,omitempty"`

	// Fully-qualified domain name to register.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	/*
	   ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.

	   When uploading a certificate, the following arguments are supported:
	*/
	RegionalCertificateArn string `json:"regionalCertificateArn,omitempty" yaml:"regionalCertificateArn,omitempty"`

	// Transport Layer Security (TLS) version + cipher suite for this DomainName. Valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`

	/*
	   Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   When referencing an AWS-managed certificate, the following arguments are supported:
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Certificate issued for the domain name being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
	CertificateBody string `json:"certificateBody,omitempty" yaml:"certificateBody,omitempty"`

	// Unique name to use when registering this certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`. Required if `certificate_arn` is not set.
	CertificateName string `json:"certificateName,omitempty" yaml:"certificateName,omitempty"`

	// Configuration block defining API endpoint information including type. See below.
	EndpointConfiguration types.Apigateway_DomainNameEndpointConfiguration `json:"endpointConfiguration,omitempty" yaml:"endpointConfiguration,omitempty"`

	// Mutual TLS authentication configuration for the domain name. See below.
	MutualTlsAuthentication types.Apigateway_DomainNameMutualTlsAuthentication `json:"mutualTlsAuthentication,omitempty" yaml:"mutualTlsAuthentication,omitempty"`
}
