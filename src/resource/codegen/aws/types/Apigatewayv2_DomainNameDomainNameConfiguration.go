package types

type Apigatewayv2_DomainNameDomainNameConfiguration struct {
	// Target domain name.
	TargetDomainName string `json:"targetDomainName,omitempty" yaml:"targetDomainName,omitempty"`

	// ARN of an AWS-managed certificate that will be used by the endpoint for the domain name. AWS Certificate Manager is the only supported source. Use the `aws.acm.Certificate` resource to configure an ACM certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Endpoint type. Valid values: `REGIONAL`.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// Amazon Route 53 Hosted Zone ID of the endpoint.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// ARN of the AWS-issued certificate used to validate custom domain ownership (when `certificate_arn` is issued via an ACM Private CA or `mutual_tls_authentication` is configured with an ACM-imported certificate.)
	OwnershipVerificationCertificateArn string `json:"ownershipVerificationCertificateArn,omitempty" yaml:"ownershipVerificationCertificateArn,omitempty"`

	// Transport Layer Security (TLS) version of the [security policy](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-custom-domain-tls-version.html) for the domain name. Valid values: `TLS_1_2`.
	SecurityPolicy string `json:"securityPolicy,omitempty" yaml:"securityPolicy,omitempty"`
}
